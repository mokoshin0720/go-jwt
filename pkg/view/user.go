package view

import "github.com/ispec-inc/sample/pkg/domain/model"

type UserName struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserName(m model.User) UserName {
	return UserName{
		ID: m.ID,
		Name: m.Name,
		Email: m.Email,
		Password: m.Password,
	}
}