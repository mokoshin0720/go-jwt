package user

import "github.com/ispec-inc/sample/pkg/domain/model"

type FindNameInput struct {
	ID int64
}

type AddNameInput struct {
	User model.User
}

type FindPasswordInput struct {
	Email string
}

type JwtLoginInput struct {
	Email    string
	Password string
}

type LoginUserInput struct {
	TokenString string
}