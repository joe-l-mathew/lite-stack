package functions

import (
	"context"

	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

func ListNetworks(cli *client.Client, ctx context.Context) ([]network.Summary, error) {
	return cli.NetworkList(ctx, network.ListOptions{})
}
