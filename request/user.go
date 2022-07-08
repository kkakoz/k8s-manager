package request

type LoginReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserAddReq struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
