package functions

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func Delete_Container(container_id string, cli *client.Client, ctx context.Context) error {
	// Call Docker API to create the network
	err := cli.ContainerRemove(ctx, container_id, container.RemoveOptions{})
	if err != nil {
		fmt.Println("Error deleting container")
	} else {
		fmt.Printf("Container deleted with ID: %s\n", container_id)

	}
	// Print the network ID of the created network
	return err
}
