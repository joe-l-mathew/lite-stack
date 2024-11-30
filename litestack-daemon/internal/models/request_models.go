package models

type NetworkRequest struct {
	NetworkName string `json:"network_name"` // This will map the "network_name" from the JSON body
	Subnet      string `json:"subnet"`       // This will map the "subnet" from the JSON body
}

type NetworkDeleteRequest struct {
	NetworkName string `json:"network_name"`
}

type ContainerCreationRequest struct {
	InstanceName string `json:"instance_name"`
	NetworkName  string `json:"network_name"`
}

type ConatinerDeleteRequest struct {
	ConatinerName string `json:"container_name"`
}
