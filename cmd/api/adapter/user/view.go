package user

import (
	"os"
	// "time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ispec-inc/sample/pkg/view"
	"github.com/ispec-inc/sample/src/user"
)

type GetNameResponse struct {
	UserName view.UserName `json:"user_name"`
}

type AddNameResponse struct {
	UserName view.UserName `json:"user_name"`
}

type GetPasswordResponse struct {
	UserPassword view.UserName `json:"user_name"`
}

type GetJWTResponse struct {
	Token string `json:"token"`
}

func CreateToken(out user.JwtLoginOutput) (string, error) {
	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = out.User.ID
	claims["name"] = out.User.Name
	claims["email"] = out.User.Email
	// claims["exp"] = time.Now().Add(time.Minute * 15).Unix() // 15分で認証が切れる

	// シークレットキーをもとにJWTを発行
	tokenString, jerr := token.SignedString([]byte(os.Getenv("SIGNINKEY")))
	return tokenString, jerr
}