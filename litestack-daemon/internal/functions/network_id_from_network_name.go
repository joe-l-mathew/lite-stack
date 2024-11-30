package functions

import (
	"litestack-daemon/internal/dockerclient"
	"log"

	"github.com/docker/docker/api/types/network"
)

func GetNetworkIdFromName(name string) string {
	networks, err := dockerclient.CLI.NetworkList(dockerclient.CTX, network.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing Docker networks: %v", err)
		return ""
	}

	// Iterate through the networks and find the one with the matching name
	for _, network := range networks {
		if network.Name == name {
			// Return the network ID
			return network.ID
		}
	}

	// If no matching network was found, return an empty string
	return ""
}
