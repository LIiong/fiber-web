package model

import (
	"fiber-web/pkg/errno"
)

type User struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

func (u *User) Login() error {

	name := u.Name
	pass := u.Password
	if name != "liion" || pass != "123456" {
		return errno.UserNotExits
	}
	return nil
}