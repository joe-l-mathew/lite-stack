package functions

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/client"
)

func Delete_Network(network_id string, cli *client.Client, ctx context.Context) {
	// Call Docker API to create the network
	err := cli.NetworkRemove(ctx, network_id)
	if err != nil {
		log.Fatalf("Error creating network: %v", err)
	}
	// Print the network ID of the created network
	fmt.Printf("Network deleted with ID: %s\n", network_id)
}
