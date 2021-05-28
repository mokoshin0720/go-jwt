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

func Parse(tokenString string) (string, error) {
	// tokenの解析
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINKEY")), nil
	})
	// tokenに関するエラーハンドリング
	if err != nil {
		return "", err 
	}

	// tokenからemailの取得 → もっといい方法がある気がする...
	email := token.Claims.(jwt.MapClaims)["User"].(map[string]interface{})["Email"]

	return email.(string), nil
}

// tokenをstringに変換するメソッド
func (t Token) String() string {
	return string(t)
}