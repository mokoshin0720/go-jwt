package user

import "github.com/ispec-inc/sample/pkg/view"

type GetNameResponse struct {
	UserName view.UserName `json:"user_name"`
}

type AddNameResponse struct {
	UserName view.UserName `json:"user_name"`
}