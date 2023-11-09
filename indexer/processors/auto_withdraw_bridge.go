package processors

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"os"
	"time"

	"github.com/ethereum-optimism/optimism/indexer/config"
	"github.com/ethereum-optimism/optimism/indexer/database"
	"github.com/ethereum-optimism/optimism/indexer/etl"
	"github.com/ethereum-optimism/optimism/indexer/node"
	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

type AutoWithdrawBridgeProcessor struct {
	log log.Logger
	db  *database.DB

	l1Etl       *etl.L1ETL
	l1EthClient node.EthClient
	l2EthClient node.EthClient
	chainConfig config.ChainConfig
}

func NewAutoWithdrawBridgeProcessor(
	log log.Logger,
	db *database.DB,
	l1Etl *etl.L1ETL,
	l1EthClient node.EthClient,
	l2EthClient node.EthClient,
	chainConfig config.ChainConfig,
) (*AutoWithdrawBridgeProcessor, error) {
	log = log.New("processor", "auto_withdraw_bridge")

	return &AutoWithdrawBridgeProcessor{log, db, l1Etl, l1EthClient, l2EthClient, chainConfig}, nil
}

func (b *AutoWithdrawBridgeProcessor) Start(ctx context.Context) error {
	done := ctx.Done()

	// Fire off independently on startup to check for
	// new data or if we've indexed new L1 data.
	l1EtlUpdates := b.l1Etl.Notify()
	startup := make(chan interface{}, 1)
	startup <- nil

	b.log.Info("starting bridge processor...")
	for {
		select {
		case <-done:
			b.log.Info("stopping bridge processor")
			return nil

		// Tickers
		case <-startup:
		case <-l1EtlUpdates:
		}

		err := b.run()
		if err != nil {
			log.Error("AutoWithdrawBridgeProcessor run error", "err", err)
		}
	}
}

func (b *AutoWithdrawBridgeProcessor) run() error {
	if unprovens, err := b.db.BridgeTransactions.L2UnprovenAutoWithdrawals(); err != nil {
		return err
	} else {
		for _, unproven := range unprovens {
			initiatedL2Event, err := b.db.ContractEvents.L2ContractEvent(unproven.InitiatedL2EventGUID)
			if err != nil {
				return err
			}
			if initiatedL2Event == nil {
				return fmt.Errorf("cannot find initiated l2 event for unproven withdrawal, initiated_l2_event_guid: %s", unproven.ProvenL1EventGUID)
			}

			// TODO bilibili
			const proposePeriod = 60 * 2
			if initiatedL2Event.Timestamp+proposePeriod < uint64(time.Now().Unix()) {
				err := b.proveWithdrawalTransaction(initiatedL2Event.TransactionHash)
				if err != nil {
					return fmt.Errorf("proveWithdrawalTransaction error: %v", err)
				}
			}
		}
	}

	if unfinalizeds, err := b.db.BridgeTransactions.L2UnfinalizedAutoWithdrawals(); err != nil {
		return err
	} else {
		for _, unfinalized := range unfinalizeds {
			initiatedL2Event, err := b.db.ContractEvents.L2ContractEvent(unfinalized.InitiatedL2EventGUID)
			if err != nil {
				return err
			}
			if initiatedL2Event == nil {
				return fmt.Errorf("cannot find initiated l2 event for unfinalized withdrawal, initiated_l2_event_guid: %s", unfinalized.ProvenL1EventGUID)
			}

			const challengePeriod = 60 * 10
			if initiatedL2Event.Timestamp+challengePeriod < uint64(time.Now().Unix()) {
				err := b.finalizeMessage(initiatedL2Event.TransactionHash)
				if err != nil {
					return fmt.Errorf("finalizeMessage error: %v", err)
				}
			}
		}
	}

	return nil
}

func (b *AutoWithdrawBridgeProcessor) proveWithdrawalTransaction(txHash common.Hash) error {
	receipt, err := b.l2EthClient.GetTransactionReceipt(txHash)
	if err != nil {
		return err
	}
	l2BlockNumber := receipt.BlockNumber

	sentMessageEvents, err := b.getSentMessagesByReceipt(receipt)
	if err != nil {
		return err
	}
	messagePassedEvents, err := b.getMessagePassedMessagesFromReceipt(receipt)
	if err != nil {
		return err
	}

	// TODO handle multicall
	sentMessageEvent := sentMessageEvents[0]
	messagePassedEvent := messagePassedEvents[0]

	withdrawalTx, err := b.toLowLevelMessage(&sentMessageEvent, &messagePassedEvent)
	if err != nil {
		return fmt.Errorf("toLowLevelMessage err: %v", err)
	}

	hash, err := b.hashWithdrawal(withdrawalTx)
	if err != nil {
		return fmt.Errorf("hashWithdrawal err: %v", err)
	}

	messageSlot, err := b.hashMessageHash(hash)
	if err != nil {
		return fmt.Errorf("hashMesaageHash err: %v", err)
	}

	l2OutputIndex, outputProposal, err := b.getL2OutputAfter(l2BlockNumber)
	if err != nil {
		return err
	}
	accountResult, err := b.l2EthClient.GetProof(
		b.chainConfig.L2Contracts.L2ToL1MessagePasser,
		[]string{"0x" + messageSlot},
		outputProposal.L2BlockNumber,
	)
	if err != nil {
		return fmt.Errorf("GetProof err: %v", err)
	}

	outputProposalBlock, err := b.l2EthClient.BlockHeaderByNumber(outputProposal.L2BlockNumber)
	if err != nil {
		return fmt.Errorf("get output proposal block error: %v", err)
	}

	withdrawalProof := accountResult.StorageProof[0]
	withdrawalProof2Bytes := make([][]byte, 0)
	for _, p1 := range withdrawalProof.Proof {
		withdrawalProof2Bytes = append(withdrawalProof2Bytes, p1)
	}

	outputRootProof := bindings.TypesOutputRootProof{
		Version:                  common.HexToHash("0x"),
		StateRoot:                outputProposalBlock.Root,
		MessagePasserStorageRoot: accountResult.StorageHash,
		LatestBlockhash:          outputProposalBlock.Hash(),
	}

	l1ChainId, err := b.l1EthClient.ChainId()
	if err != nil {
		return err
	}

	envIndexerPrivkey := os.Getenv("INDEXER_AUTO_WITHDRAW_BRIDGE_PRIVKEY")
	indexerPrivkey, err := crypto.HexToECDSA(envIndexerPrivkey)
	if err != nil {
		return err
	}

	pubKey := indexerPrivkey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	pubKeyBytes := crypto.FromECDSAPub(pubKeyECDSA)
	pubKeyHash := crypto.Keccak256(pubKeyBytes[1:])[12:]
	fromAddress := common.HexToAddress(hexutil.Encode(pubKeyHash))

	const gasPrice = 9000000000 // 9 GWei

	optimismPortalTransactor, _ := bindings.NewOptimismPortalTransactor(
		b.chainConfig.L1Contracts.OptimismPortalProxy,
		b.l1EthClient,
	)
	signedTx, err := optimismPortalTransactor.ProveWithdrawalTransaction(
		&bind.TransactOpts{
			From:     fromAddress,
			GasPrice: big.NewInt(gasPrice),
			Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
				return types.SignTx(tx, types.NewEIP155Signer(l1ChainId), indexerPrivkey)
			},
		},
		*withdrawalTx,
		l2OutputIndex,
		outputRootProof,
		withdrawalProof2Bytes,
	)
	if err != nil {
		return err
	}

	b.log.Info("ProveWithdrawalTransaction", "tx_hash", signedTx.Hash())
	return nil
}

// [finalizeMessage](https://github.com/ethereum-optimism/optimism/blob/d90e7818de894f0bc93ae7b449b9049416bda370/packages/sdk/src/cross-chain-messenger.ts#L1611)
func (b *AutoWithdrawBridgeProcessor) finalizeMessage(txHash common.Hash) error {
	receipt, err := b.l2EthClient.GetTransactionReceipt(txHash)
	if err != nil {
		return err
	}

	sentMessageEvents, err := b.getSentMessagesByReceipt(receipt)
	if err != nil {
		return err
	}
	messagePassedEvents, err := b.getMessagePassedMessagesFromReceipt(receipt)
	if err != nil {
		return err
	}

	// TODO handle multicall
	sentMessageEvent := sentMessageEvents[0]
	messagePassedEvent := messagePassedEvents[0]

	withdrawalTx, err := b.toLowLevelMessage(&sentMessageEvent, &messagePassedEvent)
	if err != nil {
		return fmt.Errorf("toLowLevelMessage err: %v", err)
	}

	l1ChainId, err := b.l1EthClient.ChainId()
	if err != nil {
		return err
	}

	envIndexerPrivkey := os.Getenv("INDEXER_AUTO_WITHDRAW_BRIDGE_PRIVKEY")
	indexerPrivkey, err := crypto.HexToECDSA(envIndexerPrivkey)
	if err != nil {
		return err
	}

	pubKey := indexerPrivkey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	pubKeyBytes := crypto.FromECDSAPub(pubKeyECDSA)
	pubKeyHash := crypto.Keccak256(pubKeyBytes[1:])[12:]
	fromAddress := common.HexToAddress(hexutil.Encode(pubKeyHash))

	const gasPrice = 9000000000 // 9 GWei

	optimismPortalTransactor, _ := bindings.NewOptimismPortalTransactor(
		b.chainConfig.L1Contracts.OptimismPortalProxy,
		b.l1EthClient,
	)
	signedTx, err := optimismPortalTransactor.FinalizeWithdrawalTransaction(
		&bind.TransactOpts{
			From:     fromAddress,
			GasPrice: big.NewInt(gasPrice),
			Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
				return types.SignTx(tx, types.NewEIP155Signer(l1ChainId), indexerPrivkey)
			},
		},
		*withdrawalTx,
	)
	if err != nil {
		return err
	}

	b.log.Info("FinalizeWithdrawalTransaction", "tx_hash", signedTx.Hash())
	return nil
}

func (b *AutoWithdrawBridgeProcessor) hashWithdrawal(w *bindings.TypesWithdrawalTransaction) (string, error) {
	uint256Type, _ := abi.NewType("uint256", "", nil)
	addressType, _ := abi.NewType("address", "", nil)
	bytesType, _ := abi.NewType("bytes", "", nil)
	types_ := abi.Arguments{
		{Type: uint256Type},
		{Type: addressType},
		{Type: addressType},
		{Type: uint256Type},
		{Type: uint256Type},
		{Type: bytesType},
	}
	encoded, err := types_.Pack(w.Nonce, w.Sender, w.Target, w.Value, w.GasLimit, w.Data)
	if err != nil {
		return "", err
	}
	result := crypto.Keccak256(encoded)
	return common.Bytes2Hex(result), nil
}

func (b *AutoWithdrawBridgeProcessor) hashMessageHash(messageHash string) (string, error) {
	uint256Type, _ := abi.NewType("uint256", "", nil)
	bytes32Type, _ := abi.NewType("bytes32", "", nil)
	types_ := abi.Arguments{
		{
			Type: bytes32Type,
		},
		{
			Type: uint256Type,
		},
	}

	encoded, err := types_.Pack(common.HexToHash(messageHash), big.NewInt(0))
	if err != nil {
		return "", err
	}

	return common.Bytes2Hex(crypto.Keccak256(encoded)), nil
}

type L2CrossDomainMessengerSentMessageExtension1 struct {
	bindings.L2CrossDomainMessengerSentMessage
	Value *big.Int
}

// getSentMessagesByReceipt retrieves all cross chain messages sent within a given transaction.
func (b *AutoWithdrawBridgeProcessor) getSentMessagesByReceipt(receipt *types.Receipt) ([]L2CrossDomainMessengerSentMessageExtension1, error) {
	L2CrossDomainMessengerAbi, _ := bindings.L2CrossDomainMessengerMetaData.GetAbi()
	addressType, _ := abi.NewType("address", "", nil)

	// Filter SentMessage(address indexed target, address sender, bytes message, uint256 messageNonce, uint256 gasLimit)
	sentMessageEvents := make([]L2CrossDomainMessengerSentMessageExtension1, 0)
	for i, l := range receipt.Logs {
		if l.Address == b.chainConfig.L2Contracts.L2CrossDomainMessenger &&
			len(l.Topics) > 0 &&
			l.Topics[0] == L2CrossDomainMessengerAbi.Events["SentMessage"].ID {

			sentMessageEvent := bindings.L2CrossDomainMessengerSentMessage{}
			err := abi.ParseTopics(
				&sentMessageEvent,
				[]abi.Argument{
					{
						Name:    "target",
						Type:    addressType,
						Indexed: true,
					},
				},
				l.Topics[1:],
			)
			if err != nil {
				return nil, fmt.Errorf("parse indexed event arguments from log.topics of SentMessage event, err: %v", err)
			}

			// NOTE: log.Data only contains the non-indexed arguments
			err = L2CrossDomainMessengerAbi.UnpackIntoInterface(&sentMessageEvent, "SentMessage", l.Data)
			if err != nil {
				return nil, fmt.Errorf("parse non-indexed event arguments from log.data of SentMessage event, err: %v", err)
			}

			if i+1 < len(receipt.Logs) &&
				receipt.Logs[i+1].Address == b.chainConfig.L2Contracts.L2CrossDomainMessenger &&
				len(receipt.Logs[i+1].Topics) > 1 &&
				receipt.Logs[i+1].Topics[0] == L2CrossDomainMessengerAbi.Events["SentMessageExtension1"].ID {

				sentMessageExtension1 := bindings.L2CrossDomainMessengerSentMessageExtension1{}
				err := L2CrossDomainMessengerAbi.UnpackIntoInterface(&sentMessageExtension1, "SentMessageExtension1", receipt.Logs[i+1].Data)
				if err != nil {
					return nil, err
				}

				sentMessageEvents = append(sentMessageEvents, L2CrossDomainMessengerSentMessageExtension1{
					L2CrossDomainMessengerSentMessage: sentMessageEvent,
					Value:                             sentMessageExtension1.Value,
				})
			}
		}
	}

	return sentMessageEvents, nil
}

func (b *AutoWithdrawBridgeProcessor) getMessagePassedMessagesFromReceipt(receipt *types.Receipt) ([]bindings.L2ToL1MessagePasserMessagePassed, error) {
	L2ToL1MessagePasserAbi, _ := bindings.L2ToL1MessagePasserMetaData.GetAbi()
	uint256Type, _ := abi.NewType("uint256", "", nil)
	addressType, _ := abi.NewType("address", "", nil)

	messagePassedLogs := make([]*types.Log, 0)
	for _, l := range receipt.Logs {
		if l.Address == b.chainConfig.L2Contracts.L2ToL1MessagePasser &&
			len(l.Topics) > 0 &&
			l.Topics[0] == L2ToL1MessagePasserAbi.Events["MessagePassed"].ID {
			messagePassedLogs = append(messagePassedLogs, l)
		}
	}
	if len(messagePassedLogs) == 0 {
		return nil, errors.New("no MessagePassed event")
	}

	// Parse SentMessage events
	messagePassedEvents := make([]bindings.L2ToL1MessagePasserMessagePassed, len(messagePassedLogs))
	for i, l := range messagePassedLogs {
		messagePassedEvent := bindings.L2ToL1MessagePasserMessagePassed{}
		err := abi.ParseTopics(
			&messagePassedEvent,
			[]abi.Argument{
				{Name: "nonce", Type: uint256Type, Indexed: true},
				{Name: "sender", Type: addressType, Indexed: true},
				{Name: "target", Type: addressType, Indexed: true},
			},
			l.Topics[1:],
		)
		if err != nil {
			return nil, fmt.Errorf("parse indexed event arguments from log.topics of MessagePassed event, err: %v", err)
		}

		// NOTE: log.Data only contains the non-indexed arguments
		err = L2ToL1MessagePasserAbi.UnpackIntoInterface(&messagePassedEvent, "MessagePassed", l.Data)
		if err != nil {
			return nil, fmt.Errorf("parse non-indexed event arguments from log.data of SentMessage event, err: %v", err)
		}

		// NOTE: log.Data only contains the non-indexed arguments
		err = L2ToL1MessagePasserAbi.UnpackIntoInterface(&messagePassedEvent, "MessagePassed", l.Data)
		if err != nil {
			return nil, fmt.Errorf("parse non-indexed event arguments from log.data of MessagePassed event, err: %v", err)
		}

		messagePassedEvents[i] = messagePassedEvent
	}

	return messagePassedEvents, nil
}

func (b *AutoWithdrawBridgeProcessor) getL2OutputAfter(l2BlockNumber *big.Int) (*big.Int, *bindings.TypesOutputProposal, error) {
	l2OutputOracleCaller, err := bindings.NewL2OutputOracleCaller(
		b.chainConfig.L1Contracts.L2OutputOracleProxy,
		b.l1EthClient,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("NewL2OutputOracleCaller err: %v", err)
	}

	// [getBedrockMessageProof](https://github.com/ethereum-optimism/optimism/blob/d90e7818de894f0bc93ae7b449b9049416bda370/packages/sdk/src/cross-chain-messenger.ts#L1916)
	l2OutputIndex, err := l2OutputOracleCaller.GetL2OutputIndexAfter(&bind.CallOpts{}, l2BlockNumber)
	if err != nil {
		// TODO handle "execution reverted: L2OutputOracle: cannot get output for a block that has not been proposed, call: 0"
		return nil, nil, fmt.Errorf("GetL2OutputIndexAfter err: %v", err)
	}

	outputProposal, err := l2OutputOracleCaller.GetL2Output(&bind.CallOpts{}, l2OutputIndex)
	if err != nil {
		return nil, nil, fmt.Errorf("GetL2Output err: %v", err)
	}

	return l2OutputIndex, &outputProposal, nil
}

func (b *AutoWithdrawBridgeProcessor) toLowLevelMessage(
	sentMessageEvent *L2CrossDomainMessengerSentMessageExtension1,
	messagePassedEvent *bindings.L2ToL1MessagePasserMessagePassed,
) (*bindings.TypesWithdrawalTransaction, error) {
	// Encode "relayMessage" with signature, the result will be attached to [WithdrawalTransaction.Data](https://github.com/ethereum-optimism/optimism/blob/f54a2234f2f350795552011f35f704a3feb56a08/packages/contracts-bedrock/src/libraries/Types.sol#L68)
	L2CrossDomainMessengerAbi, _ := bindings.L2CrossDomainMessengerMetaData.GetAbi()
	relayMessageCalldata, err := L2CrossDomainMessengerAbi.Pack(
		"relayMessage",
		sentMessageEvent.MessageNonce,
		sentMessageEvent.Sender,
		sentMessageEvent.Target,
		sentMessageEvent.Value,
		sentMessageEvent.GasLimit,
		sentMessageEvent.Message,
	)
	if err != nil {
		return nil, fmt.Errorf("encode relayMessage calldata, err: %v", err)
	}

	withdrawalTx := bindings.TypesWithdrawalTransaction{
		Nonce:    messagePassedEvent.Nonce,
		Sender:   b.chainConfig.L2Contracts.L2CrossDomainMessenger,
		Target:   b.chainConfig.L1Contracts.L1CrossDomainMessengerProxy,
		Value:    sentMessageEvent.Value,
		GasLimit: messagePassedEvent.GasLimit,
		Data:     relayMessageCalldata,
	}
	return &withdrawalTx, nil
}
