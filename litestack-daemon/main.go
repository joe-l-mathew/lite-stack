package main

import (
	"litestack-daemon/internal/config"
	"litestack-daemon/internal/dockerclient"
)

func main() {
	// Get instance of cli and ctx
	cli, ctx := dockerclient.CreateClient()
	// close the connction on exit
	defer cli.Close()
	// Init will ensure the default and public network
	//subnets are available in the installed environment
	config.InitEnvironment(cli, ctx)
}
