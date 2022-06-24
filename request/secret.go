package request

type SecretAddReq struct {
	Namespace string            `json:"namespace"`
	Data      map[string]string `json:"data"`
}
