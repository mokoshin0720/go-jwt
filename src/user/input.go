package user

import "github.com/ispec-inc/sample/pkg/domain/model"

type FindNameInput struct {
	ID int64
}

type AddNameInput struct {
	User model.User
}