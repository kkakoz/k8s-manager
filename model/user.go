package model

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Salt     string `json:"-"`
}

func (u *User) Format() map[string]any {
	return map[string]any{
		"id":   u.ID,
		"name": u.Name,
	}
}
