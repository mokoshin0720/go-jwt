package view

import "github.com/ispec-inc/sample/pkg/domain/model"

type UserName struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

func NewUserName(m model.User) UserName {
	return UserName{
		ID: m.ID,
		Name: m.Name,
	}
}