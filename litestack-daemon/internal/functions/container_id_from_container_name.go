package functions

import (
	"litestack-daemon/internal/dockerclient"
	"log"

	"github.com/docker/docker/api/types/container"
)

func GetContainerIdFromName(name string) string {
	// List all containers (including stopped ones)
	containers, err := dockerclient.CLI.ContainerList(dockerclient.CTX, container.ListOptions{All: true})
	if err != nil {
		log.Fatalf("Error listing Docker containers: %v", err)
		return ""
	}

	// Iterate through the containers and find the one with the matching name
	for _, container := range containers {
		// Checking if the container name matches
		for _, containerName := range container.Names {
			// Strip the leading '/' from the container name as Docker adds it to the names
			if containerName == "/"+name {
				// Return the container ID
				return container.ID
			}
		}
	}

	// If no matching container was found, return an empty string
	return ""
}
