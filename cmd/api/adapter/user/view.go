package user

import (
	"github.com/ispec-inc/sample/pkg/view"
)

type GetNameResponse struct {
	UserName view.UserName `json:"user"`
}

type AddNameResponse struct {
	UserName view.UserName `json:"user"`
}

type GetPasswordResponse struct {
	UserPassword view.UserName `json:"user"`
}

type GetJWTResponse struct {
	Token string `json:"token"`
}