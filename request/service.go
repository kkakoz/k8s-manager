package request

type ServiceAddReq struct {
	Name     string `json:"name"`
	Selector string `json:"selector"`
	Type     uint8  `json:"type"` // 1为cluster ip 2为note pod
	Port     int32  `json:"port"`
	NodePort int32  `json:"node_port"`
}
