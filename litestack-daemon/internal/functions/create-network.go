package functions

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

// Create_Networks creates a Docker network with the given name
func Create_Networks(name string, cli *client.Client, ctx context.Context) {
	// Set up network creation options
	networkOptions := network.CreateOptions{
		Driver: "bridge",
	}
	// Call Docker API to create the network
	networkResponse, err := cli.NetworkCreate(ctx, name, networkOptions)
	if err != nil {
		log.Fatalf("Error creating network: %v", err)
	}
	// Print the network ID of the created network
	fmt.Printf("Network created with ID: %s\n", networkResponse.ID)
}
