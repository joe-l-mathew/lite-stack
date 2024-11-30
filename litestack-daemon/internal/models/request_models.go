package request_models

type NetworkRequest struct {
	NetworkName string `json:"network_name"` // This will map the "network_name" from the JSON body
	Subnet      string `json:"subnet"`       // This will map the "subnet" from the JSON body
}
