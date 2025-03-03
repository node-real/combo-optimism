package sources

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/eth/catalyst"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/sources/caching"
)

type EngineClientConfig struct {
	L2ClientConfig
}

func EngineClientDefaultConfig(config *rollup.Config) *EngineClientConfig {
	return &EngineClientConfig{
		// engine is trusted, no need to recompute responses etc.
		L2ClientConfig: *L2ClientDefaultConfig(config, true),
	}
}

// EngineClient extends L2Client with engine API bindings.
type EngineClient struct {
	*L2Client
	*EngineAPIClient
}

func NewEngineClient(client client.RPC, log log.Logger, metrics caching.Metrics, config *EngineClientConfig) (*EngineClient, error) {
	l2Client, err := NewL2Client(client, log, metrics, &config.L2ClientConfig)
	if err != nil {
		return nil, err
	}

	engineAPIClient := NewEngineAPIClient(client, log, config.RollupCfg)

	return &EngineClient{
		L2Client:        l2Client,
		EngineAPIClient: engineAPIClient,
	}, nil
}

// EngineAPIClient is an RPC client for the Engine API functions.
type EngineAPIClient struct {
	RPC client.RPC
	log log.Logger
	evp EngineVersionProvider
}

type EngineVersionProvider interface {
	ForkchoiceUpdatedVersion(attr *eth.PayloadAttributes) eth.EngineAPIMethod
	NewPayloadVersion(timestamp uint64) eth.EngineAPIMethod
	GetPayloadVersion(timestamp uint64) eth.EngineAPIMethod
	SealPayloadVersion(timestamp uint64) eth.EngineAPIMethod
}

func NewEngineAPIClient(rpc client.RPC, l log.Logger, evp EngineVersionProvider) *EngineAPIClient {
	return &EngineAPIClient{
		RPC: rpc,
		log: l,
		evp: evp,
	}
}

// EngineVersionProvider returns the underlying engine version provider used for
// resolving the correct Engine API versions.
func (s *EngineAPIClient) EngineVersionProvider() EngineVersionProvider { return s.evp }

// ForkchoiceUpdate updates the forkchoice on the execution client. If attributes is not nil, the engine client will also begin building a block
// based on attributes after the new head block and return the payload ID.
//
// The RPC may return three types of errors:
// 1. Processing error: ForkchoiceUpdatedResult.PayloadStatusV1.ValidationError or other non-success PayloadStatusV1,
// 2. `error` as eth.InputError: the forkchoice state or attributes are not valid.
// 3. Other types of `error`: temporary RPC errors, like timeouts.
func (s *EngineAPIClient) ForkchoiceUpdate(ctx context.Context, fc *eth.ForkchoiceState, attributes *eth.PayloadAttributes) (*eth.ForkchoiceUpdatedResult, error) {
	llog := s.log.New("state", fc)       // local logger
	tlog := llog.New("attr", attributes) // trace logger
	tlog.Trace("Sharing forkchoice-updated signal")
	fcCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	var result eth.ForkchoiceUpdatedResult
	method := s.evp.ForkchoiceUpdatedVersion(attributes)
	err := s.RPC.CallContext(fcCtx, &result, string(method), fc, attributes)
	if err == nil {
		tlog.Trace("Shared forkchoice-updated signal")
		if attributes != nil { // block building is optional, we only get a payload ID if we are building a block
			tlog.Trace("Received payload id", "payloadId", result.PayloadID)
		}
		return &result, nil
	} else {
		llog.Warn("Failed to share forkchoice-updated signal", "err", err)
		if rpcErr, ok := err.(rpc.Error); ok {
			code := eth.ErrorCode(rpcErr.ErrorCode())
			switch code {
			case eth.InvalidForkchoiceState, eth.InvalidPayloadAttributes:
				return nil, eth.InputError{
					Inner: err,
					Code:  code,
				}
			default:
				return nil, fmt.Errorf("unrecognized rpc error: %w", err)
			}
		}
		return nil, err
	}
}

// NewPayload executes a full block on the execution engine.
// This returns a PayloadStatusV1 which encodes any validation/processing error,
// and this type of error is kept separate from the returned `error` used for RPC errors, like timeouts.
func (s *EngineAPIClient) NewPayload(ctx context.Context, payload *eth.ExecutionPayload, parentBeaconBlockRoot *common.Hash) (*eth.PayloadStatusV1, error) {
	e := s.log.New("block_hash", payload.BlockHash)
	e.Trace("sending payload for execution")

	execCtx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	var result eth.PayloadStatusV1

	var err error
	switch method := s.evp.NewPayloadVersion(uint64(payload.Timestamp)); method {
	case eth.NewPayloadV3:
		err = s.RPC.CallContext(execCtx, &result, string(method), payload, []common.Hash{}, parentBeaconBlockRoot)
	case eth.NewPayloadV2:
		err = s.RPC.CallContext(execCtx, &result, string(method), payload)
	default:
		return nil, fmt.Errorf("unsupported NewPayload version: %s", method)
	}

	e.Trace("Received payload execution result", "status", result.Status, "latestValidHash", result.LatestValidHash, "message", result.ValidationError)
	if err != nil {
		if strings.Contains(err.Error(), derive.ErrELSyncTriggerUnexpected.Error()) {
			result.Status = eth.ExecutionSyncing
			return &result, err
		}
		e.Error("Payload execution failed", "err", err)
		return nil, fmt.Errorf("failed to execute payload: %w", err)
	}
	return &result, nil
}

// GetPayload gets the execution payload associated with the PayloadId.
// There may be two types of error:
// 1. `error` as eth.InputError: the payload ID may be unknown
// 2. Other types of `error`: temporary RPC errors, like timeouts.
func (s *EngineAPIClient) GetPayload(ctx context.Context, payloadInfo eth.PayloadInfo) (*eth.ExecutionPayloadEnvelope, error) {
	e := s.log.New("payload_id", payloadInfo.ID)
	e.Trace("getting payload")
	var result eth.ExecutionPayloadEnvelope
	method := s.evp.GetPayloadVersion(payloadInfo.Timestamp)
	err := s.RPC.CallContext(ctx, &result, string(method), payloadInfo.ID)
	if err != nil {
		e.Warn("Failed to get payload", "payload_id", payloadInfo.ID, "err", err)
		if rpcErr, ok := err.(rpc.Error); ok {
			code := eth.ErrorCode(rpcErr.ErrorCode())
			switch code {
			case eth.UnknownPayload:
				return nil, eth.InputError{
					Inner: err,
					Code:  code,
				}
			default:
				return nil, fmt.Errorf("unrecognized rpc error: %w", err)
			}
		}
		return nil, err
	}
	e.Trace("Received payload")
	return &result, nil
}

// SealPayload is a combined call of GetPayload, NewPayload, ForkchoiceUpdated via engine_opSealPayload API
func (s *EngineAPIClient) SealPayload(ctx context.Context, payloadInfo eth.PayloadInfo, fc *eth.ForkchoiceState, needPayload bool) (*eth.SealPayloadResponse, string, error) {
	e := s.log.New("payload_id", payloadInfo.ID)
	e.Trace("sealing payload")
	sCtx, sCancel := context.WithTimeout(ctx, time.Second*10)
	defer sCancel()
	var result eth.SealPayloadResponse
	method := s.evp.SealPayloadVersion(payloadInfo.Timestamp)
	err := s.RPC.CallContext(sCtx, &result, string(method), payloadInfo.ID, fc, needPayload)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), eth.GetPayloadStage):
			e.Error("Seal payload get payload stage failed", "payload_id", payloadInfo.ID, "err", err)
			if rpcErr, ok := err.(rpc.Error); ok {
				code := eth.ErrorCode(rpcErr.ErrorCode())
				e.Error("Seal payload get payload stage failed", "err code", code, "payload_id", payloadInfo.ID, "err", err)
				switch code {
				case eth.UnknownPayload:
					return nil, eth.GetPayloadStage, eth.InputError{
						Inner: err,
						Code:  code,
					}
				default:
					return nil, eth.GetPayloadStage, fmt.Errorf("seal payload get payload stage unrecognized rpc error: %w", err)
				}
			}
			return nil, eth.GetPayloadStage, err
		case strings.Contains(err.Error(), eth.NewPayloadStage):
			e.Error("Seal payload new payload stage execution failed", "payload_id", payloadInfo.ID, "err", err)
			if strings.Contains(err.Error(), derive.ErrELSyncTriggerUnexpected.Error()) {
				result.PayloadStatus.Status = eth.ExecutionSyncing
				return &result, eth.NewPayloadStage, err
			}
			return nil, eth.NewPayloadStage, fmt.Errorf("seal payload failed to execute payload: %w", err)
		case strings.Contains(err.Error(), eth.ForkchoiceUpdatedStage):
			e.Error("Seal payload forkchoice updated stage failed to share forkchoice-updated signal", "payload_id", payloadInfo.ID, "err", err)
			if rpcErr, ok := err.(rpc.Error); ok {
				code := eth.ErrorCode(rpcErr.ErrorCode())
				switch code {
				case eth.InvalidForkchoiceState, eth.InvalidPayloadAttributes:
					return nil, eth.ForkchoiceUpdatedStage, eth.InputError{
						Inner: err,
						Code:  code,
					}
				default:
					return nil, eth.ForkchoiceUpdatedStage, fmt.Errorf("seal payload forkchoice updated stage unrecognized rpc error: %w", err)
				}
			}
			return nil, eth.ForkchoiceUpdatedStage, err
		default:
			e.Error("Seal payload network stage failed", "payload_id", payloadInfo.ID, "err", err)
			return nil, result.ErrStage, err
		}
	}
	e.Trace("Sealed payload")
	return &result, result.ErrStage, nil
}

func (s *EngineAPIClient) SignalSuperchainV1(ctx context.Context, recommended, required params.ProtocolVersion) (params.ProtocolVersion, error) {
	var result params.ProtocolVersion
	err := s.RPC.CallContext(ctx, &result, "engine_signalSuperchainV1", &catalyst.SuperchainSignal{
		Recommended: recommended,
		Required:    required,
	})
	return result, err
}
