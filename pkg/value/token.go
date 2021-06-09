package value

import (
	"fmt"
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

func Parse(tokenString string) (*Claims, error) {
	// tokenの解析
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SIGNINKEY")), nil
	})
	// tokenに関するエラーハンドリング
	if err != nil {
		return nil, err 
	}

	claims, ok := token.Claims.(*Claims)
	if !(ok && token.Valid) {
		return nil, fmt.Errorf(("invalid Claims"))
	}

	return claims, nil
}

// tokenをstringに変換するメソッド
func (t Token) String() string {
	return string(t)
}