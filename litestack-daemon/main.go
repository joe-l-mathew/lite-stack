package main

import (
	"litestack-daemon/api/handlers"
	"litestack-daemon/internal/config"
	"litestack-daemon/internal/dockerclient"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Get instance of cli and ctx
	cli, ctx := dockerclient.CreateClient()
	// close the connction on exit
	defer cli.Close()
	// Init will ensure the default and public network
	//subnets are available in the installed environment
	config.InitEnvironment(cli, ctx)
	router := mux.NewRouter()
	handlers.NetworkHandler(router)
	log.Println("Starting the API server on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
