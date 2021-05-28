package value

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Token string

func NewToken(claims Claims) (Token, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tkn.SignedString([]byte(os.Getenv("SIGNINKEY")))
	if err != nil {
		return "", err
	}
	return Token(tokenString), nil
}

// tokenをstringに変換するメソッド
func (t Token) String() string {
	return string(t)
}