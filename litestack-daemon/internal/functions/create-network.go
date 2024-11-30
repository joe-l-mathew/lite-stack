package functions

import (
	"context"

	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

// Create_Networks creates a Docker network with the given name
func Create_Networks(name string, cli *client.Client, ctx context.Context, subnet string) (network.CreateResponse, error) {
	// Set up the IPAM configuration
	var ipamConfig *network.IPAM
	if subnet != "" {
		// If a subnet is provided, set up IPAM with the given subnet
		ipamConfig = &network.IPAM{
			Driver: "default", // The default IPAM driver is used by most Docker networks
			Config: []network.IPAMConfig{
				{
					Subnet: subnet, // Use the custom subnet
				},
			},
		}
	} else {
		// If no subnet is provided, use default IPAM settings
		ipamConfig = &network.IPAM{
			Driver: "default", // Docker will assign a range automatically
		}
	}

	// Set up network creation options
	networkOptions := network.CreateOptions{
		Driver: "bridge",   // Set the network driver to "bridge"
		IPAM:   ipamConfig, // Use the custom or default IPAM configuration
	}

	// Call Docker API to create the network
	return cli.NetworkCreate(ctx, name, networkOptions)

}
