package batcher

import (
	"context"
	"strings"

	"github.com/ethereum-optimism/optimism/op-node/client"
	"github.com/ethereum-optimism/optimism/op-node/sources"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// dialEthClientWithTimeout attempts to dial the L1 provider using the provided
// URL. If the dial doesn't complete within defaultDialTimeout seconds, this
// method will return an error.
func dialEthClientWithTimeout(ctx context.Context, url string) (*ethclient.Client, error) {

	ctxt, cancel := context.WithTimeout(ctx, defaultDialTimeout)
	defer cancel()

	return ethclient.DialContext(ctxt, url)
}

func dialEthClientsWithTimeout(ctx context.Context, url string) ([]*ethclient.Client, error) {

	urls := strings.Split(url, ";")

	clients := make([]*ethclient.Client, 0, len(urls))
	for _, url := range urls {
		client, err := dialEthClientWithTimeout(ctx, url)
		if err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	return clients, nil
}

// dialRollupClientWithTimeout attempts to dial the RPC provider using the provided
// URL. If the dial doesn't complete within defaultDialTimeout seconds, this
// method will return an error.
func dialRollupClientWithTimeout(ctx context.Context, url string) (*sources.RollupClient, error) {
	ctxt, cancel := context.WithTimeout(ctx, defaultDialTimeout)
	defer cancel()

	rpcCl, err := rpc.DialContext(ctxt, url)
	if err != nil {
		return nil, err
	}

	return sources.NewRollupClient(client.NewBaseRPCClient(rpcCl)), nil
}

func dialRollupClientsWithTimeout(ctx context.Context, url string) (*sources.RollupClients, error) {
	urls := strings.Split(url, ";")

	clients := make([]*sources.RollupClient, 0, len(urls))
	for _, url := range urls {
		client, err := dialRollupClientWithTimeout(ctx, url)
		if err != nil {
			return nil, err
		}

		clients = append(clients, client)
	}

	return sources.NewRollupClients(clients), nil
}
