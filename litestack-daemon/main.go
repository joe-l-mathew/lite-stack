package main

import (
	"litestack-daemon/internal/dockerclient"
	"litestack-daemon/internal/functions"
)

func main() {
	cli, ctx := dockerclient.CreateClient()
	defer cli.Close()
	// functions.Create_Networks("test_network", cli, ctx)
	functions.Delete_Network("53c0b78be3ea68d96fbb238f07175119f99485f819c20af533a14eb978d8f187", cli, ctx)
}
