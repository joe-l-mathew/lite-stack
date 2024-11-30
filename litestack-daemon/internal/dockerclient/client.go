package dockerclient

import (
	"context"
	"log"

	"github.com/docker/docker/client"
)

// CreateClient initializes and returns a Docker client
func CreateClient() (*client.Client, context.Context) {
	// Create the context
	ctx := context.Background()

	// Create the Docker client instance
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error creating Docker client: %v", err)
	}

	return cli, ctx
}
