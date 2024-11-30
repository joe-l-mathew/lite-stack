package functions

import (
	"context"
	"litestack-daemon/internal/models"
	"log"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func CreateConatainer(name string, network string, cli *client.Client, ctx context.Context) (models.IpAddressModel, error) {
	name = "litestack-" + name
	if network == "" {
		network = "litestack-default-private-net"
	} else {
		network = "litestack-" + network
	}
	containerConfig := &container.Config{
		Image: "joelmathew357/litestack-ubuntu:latest", // Specify the Docker image for the container
	}
	hostConfig := &container.HostConfig{
		NetworkMode: container.NetworkMode("litestack-public-net"),
	}
	containerCreateResp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, name)
	if err != nil {
		log.Fatalf("Error creating container: %v", err)
		return models.IpAddressModel{}, err
	}
	networkId := GetNetworkIdFromName(network)
	err = cli.NetworkConnect(ctx, networkId, containerCreateResp.ID, nil)
	if err != nil {
		log.Fatalf("Error connecting container to second network: %v", err)
		return models.IpAddressModel{}, err
	}

	// Start the container in detached mode
	err = cli.ContainerStart(ctx, containerCreateResp.ID, container.StartOptions{})
	if err != nil {
		log.Fatalf("Error starting container: %v", err)
		return models.IpAddressModel{}, err
	}
	return GetContainerIP(containerCreateResp.ID, cli, ctx)

}
