package request

type PodAddReq struct {
	Namespace  string            `json:"namespace"`
	Name       string            `json:"name"`
	Containers []Container       `json:"containers"`
	Labels     map[string]string `json:"labels"`
}

type PodLogReq struct {
	Namespace     string `form:"namespace"`
	Name          string `uri:"name"`
	ContainerName string `form:"container_name"`
	TailLien      int64  `form:"tail_lien"`
}

type PodTerminalReq struct {
	Namespace     string `form:"namespace"`
	Name          string `uri:"name"`
	ContainerName string `form:"container_name"`
}
