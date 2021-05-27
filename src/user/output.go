package user

import "github.com/ispec-inc/sample/pkg/domain/model"

type FindNameOutput struct {
	User model.User
}

type AddNameOutput struct {
	User model.User
}

type FindPasswordOutput struct {
	User model.User
}

type JwtLoginOutput struct {
	User model.User
}

type LoginUserOutput struct {
	User model.User
}