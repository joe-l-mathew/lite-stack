package handlers

import (
	"encoding/json"
	"fmt"
	"litestack-daemon/internal/dockerclient"
	"litestack-daemon/internal/functions"
	request_models "litestack-daemon/internal/models"
	"net/http"

	"github.com/gorilla/mux"
)

func ContainerHandler(router *mux.Router) {
	router.HandleFunc("/create/conatiner", createConatainer).Methods("POST")
	router.HandleFunc("/delete/conatiner", deleteConatainer).Methods("POST")
}

type ConatainerCreationSuccessResponse struct {
	Message     string   `json:"message"`
	PrivateIp   []string `json:"private_ip"`
	PublicIp    string   `json:"public_ip"`
	NetworkUsed string   `json:"newtwork"`
}
type ConatinerDeletionSuccessResponse struct {
	Message     string `json:"message"`
	ContainerId string `json:"conatainer_id"`
}

func createConatainer(w http.ResponseWriter, r *http.Request) {
	var containerReq request_models.ContainerCreationRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&containerReq)
	if err != nil {
		// If there is an error decoding, send a bad request response
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	networkId := functions.GetNetworkIdFromName("litestack-" + containerReq.NetworkName)
	if networkId == "" {
		http.Error(w, "No network found", http.StatusBadRequest)
		return
	}
	ipAddrModel, err := functions.CreateConatainer(containerReq.InstanceName,
		containerReq.NetworkName, dockerclient.CLI, dockerclient.CTX)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error creating container", http.StatusBadRequest)
		return
	}
	response := ConatainerCreationSuccessResponse{
		Message:     "Conatiner created succesfully",
		PrivateIp:   ipAddrModel.PrivateIps,
		PublicIp:    ipAddrModel.PublicIp,
		NetworkUsed: containerReq.NetworkName,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func deleteConatainer(w http.ResponseWriter, r *http.Request) {
	var containerDelReq request_models.ConatinerDeleteRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&containerDelReq)
	if err != nil {
		// If there is an error decoding, send a bad request response
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if containerDelReq.ConatinerName == "" {
		http.Error(w, "Please pass a Container as string", http.StatusBadRequest)
		return
	}
	container_id := functions.GetContainerIdFromName("litestack-" + containerDelReq.ConatinerName)
	if container_id == "" {
		http.Error(w, "Container not found", http.StatusBadRequest)
		return
	}
	err = functions.Delete_Container(container_id, dockerclient.CLI, dockerclient.CTX)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	successResponse := ConatinerDeletionSuccessResponse{
		Message:     "Container deletes successfully",
		ContainerId: container_id, // Include the network ID
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(successResponse)

}
