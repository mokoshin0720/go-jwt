package password

import "golang.org/x/crypto/bcrypt"

// 生パスワードからハッシュパスワードを生成する
func Encrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// 入力されたパスワードをハッシュ化してパスワード比較
func Compare(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}