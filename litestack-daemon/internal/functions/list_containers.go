package functions

import (
	"fmt"
	"litestack-daemon/internal/dockerclient"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

type ListContainersResponse struct {
	Name string `json:"instance_name"`
}

func ListContainers() ([]types.Container, error) {
	var containerList []types.Container
	containers, err := dockerclient.CLI.ContainerList(dockerclient.CTX, container.ListOptions{All: true})
	if err != nil {
		return []types.Container{}, err
	}
	for _, container := range containers {
		if strings.HasPrefix(container.Names[0], "/litestack-") {
			containerList = append(containerList, container)
		}
	}
	fmt.Println(containerList)
	return containerList, nil
}
