package value

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/ispec-inc/sample/pkg/domain/model"
)

type Claims struct {
	// ID int64 `json:"id"`
	User model.User
	jwt.StandardClaims
}

func NewClaims(mu model.User) Claims {
	return Claims{
		User: mu,
		StandardClaims: jwt.StandardClaims{},
	}
}