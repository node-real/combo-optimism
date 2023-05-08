package sources

import (
	"context"
	"errors"

	"golang.org/x/sync/errgroup"

	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
)

type RollupClients struct {
	Rpcs []*RollupClient
}

// NewRollupClients only used in op-batcher and op-proposer
func NewRollupClients(rpcs []*RollupClient) *RollupClients {
	return &RollupClients{
		Rpcs: rpcs,
	}
}

func (r *RollupClients) OutputAtBlock(ctx context.Context, blockNum uint64) (*eth.OutputResponse, error) {
	if len(r.Rpcs) != 0 {
		return r.Rpcs[0].OutputAtBlock(ctx, blockNum)
	}
	return nil, errors.New("not implemented")
}

func (r *RollupClients) SyncStatus(ctx context.Context) (*eth.SyncStatus, error) {
	syncStatus := func(ctx context.Context) ([]*eth.SyncStatus, error) {
		g := errgroup.Group{}

		results := make([]*eth.SyncStatus, len(r.Rpcs))
		for i, rpc := range r.Rpcs {
			i, rpc := i, rpc // https://golang.org/doc/faq#closures_and_goroutines
			g.Go(func() error {
				result, err := rpc.SyncStatus(ctx)
				if err == nil {
					results[i] = result
				}
				return err
			})
		}
		if err := g.Wait(); err != nil {
			return results, err
		}
		return results, nil
	}

	results, err := syncStatus(ctx)

	var output *eth.SyncStatus
	index := 0
	for i, result := range results {
		if result != nil {
			if output == nil || result.UnsafeL2.Number > output.UnsafeL2.Number {
				index = i
				output = result
			}
		}
	}

	if output == nil {
		return nil, err
	}
	output.HightestIndex = int64(index)
	return output, nil
}

func (r *RollupClients) RollupConfig(ctx context.Context) (*rollup.Config, error) {
	if len(r.Rpcs) != 0 {
		return r.Rpcs[0].RollupConfig(ctx)
	}
	return nil, errors.New("not implemented")
}

func (r *RollupClients) Version(ctx context.Context) (string, error) {
	if len(r.Rpcs) != 0 {
		return r.Rpcs[0].Version(ctx)
	}
	return "", errors.New("not implemented")
}
