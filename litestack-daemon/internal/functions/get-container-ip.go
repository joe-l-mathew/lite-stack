package functions

import (
	"context"
	"fmt"
	"litestack-daemon/internal/models"

	"github.com/docker/docker/client"
)

func GetContainerIP(containerID string, cli *client.Client, ctx context.Context) (models.IpAddressModel, error) {
	var privateIps []string
	// Inspect the container to get detailed information
	containerInfo, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return models.IpAddressModel{}, fmt.Errorf("failed to inspect container: %v", err)
	}

	// Get the IP address from the container's network settings
	// Assuming the container is connected to the default bridge network
	networks := containerInfo.NetworkSettings.Networks
	publicip := networks["litestack-public-net"].IPAddress
	for _, val := range networks {
		if val.IPAddress != publicip {
			privateIps = append(privateIps, val.IPAddress)
		}
	}
	ipAddressModel := models.IpAddressModel{
		PublicIp:   publicip,
		PrivateIps: privateIps,
	}
	return ipAddressModel, nil
}
