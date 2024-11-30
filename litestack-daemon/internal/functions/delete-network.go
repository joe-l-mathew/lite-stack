package functions

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

func Delete_Network(network_id string, cli *client.Client, ctx context.Context) error {
	// Call Docker API to create the network
	err := cli.NetworkRemove(ctx, network_id)
	if err != nil {
		fmt.Println("Error deleting network")
	} else {
		fmt.Printf("Network deleted with ID: %s\n", network_id)

	}
	// Print the network ID of the created network
	return err
}
