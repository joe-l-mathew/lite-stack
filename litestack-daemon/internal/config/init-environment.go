package config

import (
	"context"
	"fmt"
	"litestack-daemon/internal/functions"

	"github.com/docker/docker/client"
)

func InitEnvironment(cli *client.Client, ctx context.Context) bool {
	// Create a public network with name litestack-public
	// which will be connected with all the spinning containers
	fmt.Println("initializing environment.....")
	res, err := functions.Create_Networks("litestack-public-net", cli, ctx, "10.222.0.0/16")
	if err != nil {
		if err.Error() == "Error response from daemon: network with name litestack-public-net already exists" {
			fmt.Println("Public Network already exists Skipping")
		} else {
			fmt.Println("Error Creating Public Network")
			return false
		}
	} else {
		fmt.Println("Created Public Network with id: " + res.ID)
	}

	res, err = functions.Create_Networks("litestack-default-private-net", cli, ctx, "172.222.0.0/16")
	if err != nil {
		if err.Error() == "Error response from daemon: network with name litestack-default-private-net already exists" {
			fmt.Println("Default Private Network already exists Skipping")
		} else {
			fmt.Println("Error Creating Private Network")
			return false
		}
	} else {
		fmt.Println("Created Default Private Network with id: " + res.ID)
	}

	return true
}
