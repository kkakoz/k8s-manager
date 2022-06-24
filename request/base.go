package request

type ListReq struct {
	Label    string `query:"label"`
	Field    string `query:"field"`
	Limit    int64  `query:"limit"`
	Continue string `query:"continue"`
}

type DeleteReq struct {
	Name      string `uri:"name"`
	Namespace string `form:"namespace"`
}

type ApplyReq struct {
	Content   string `json:"content"`
	Namespace string `json:"namespace"`
}

type Container struct {
	Name          string `json:"name"`
	Image         string `json:"image"`
	HostPort      int32  `json:"host_port"`
	ContainerPort int32  `json:"container_port"`
}
