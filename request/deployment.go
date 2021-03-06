package request

import corev1 "k8s.io/api/core/v1"

type DeploymentAddReq struct {
	Name       string             `json:"name"`
	Namespace  string             `json:"namespace"`
	Labels     string             `json:"labels"`
	Replicas   int32              `json:"replicas"`
	Containers []corev1.Container `json:"containers"`
}

type DeploymentRestartReq struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
