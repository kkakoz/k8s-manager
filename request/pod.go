package request

import (
	corev1 "k8s.io/api/core/v1"
)

type PodAddReq struct {
	Namespace  string             `json:"namespace"`
	Name       string             `json:"name"`
	Containers []corev1.Container `json:"containers"`
	Labels     map[string]string  `json:"labels"`
}

type PodLogReq struct {
	Namespace     string `form:"namespace"`
	Name          string `uri:"name"`
	ContainerName string `form:"container_name"`
	TailLien      int64  `form:"tail_lien"`
}
