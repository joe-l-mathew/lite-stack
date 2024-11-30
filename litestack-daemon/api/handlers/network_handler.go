package handlers

import (
	"encoding/json"
	"litestack-daemon/internal/dockerclient"
	"litestack-daemon/internal/functions"
	request_models "litestack-daemon/internal/models"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type NetworkCreationSuccessResponse struct {
	Message   string `json:"message"`
	NetworkID string `json:"network_id"`
}

func NetworkHandler(router *mux.Router) {
	router.HandleFunc("/create/network", createNewtowork).Methods("POST")
}

func createNewtowork(w http.ResponseWriter, r *http.Request) {
	var networkReq request_models.NetworkRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&networkReq)
	if err != nil {
		// If there is an error decoding, send a bad request response
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if networkReq.NetworkName == "" {
		http.Error(w, "Please pass a network_name as string", http.StatusBadRequest)
		return
	}
	res, err := functions.Create_Networks(networkReq.NetworkName, dockerclient.CLI, dockerclient.CTX, networkReq.Subnet)
	if err != nil && strings.Contains(err.Error(), "already exists") {
		http.Error(w, "Network already exist", http.StatusBadRequest)
		return
	}
	successResponse := NetworkCreationSuccessResponse{
		Message:   "Network created successfully",
		NetworkID: res.ID, // Include the network ID
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(successResponse)
}
