package functions

import (
	"litestack-daemon/internal/dockerclient"
	"strings"

	"github.com/docker/docker/api/types/network"
)

func ListNetworks() ([]network.Summary, error) {
	litestackNetworks := []network.Inspect{}
	networkList, err := dockerclient.CLI.NetworkList(dockerclient.CTX, network.ListOptions{})
	if err != nil {
		return []network.Inspect{}, err
	}
	for _, network := range networkList {
		if strings.HasPrefix(network.Name, "litestack-") {
			litestackNetworks = append(litestackNetworks, network)
		}
	}
	return litestackNetworks, nil
}
