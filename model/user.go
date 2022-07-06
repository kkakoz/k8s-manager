package model

type User struct {
	Name     string `json:"name"`
	Password string `json:"-"`
	Salt     string `json:"salt"`
}
