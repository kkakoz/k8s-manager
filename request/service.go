package request

type ServiceAddReq struct {
	Name     string            `json:"name"`
	Selector map[string]string `json:"selector"`
	Type     uint8             `json:"type"`
	Port     int32             `json:"port"`
	NodePort int32             `json:"node_port"`
}
