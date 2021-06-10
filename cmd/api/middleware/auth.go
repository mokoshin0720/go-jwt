package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/ispec-inc/sample/pkg/presenter"
	"github.com/ispec-inc/sample/pkg/registry"
	"github.com/ispec-inc/sample/pkg/value"
	"github.com/ispec-inc/sample/src/user"
)

type Auth struct {
	usecase user.Usecase
}

func NewAuth(repo registry.Repository) Auth {
	usecase := user.NewUsecase(repo)
	return Auth{usecase}
}

func (m Auth) VerifyToken(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request)  {
		ctx := r.Context()

		// requestヘッダからAuthorization情報を（Bearer込みで）取得
		token := r.Header.Get("Authorization")
		// 正しいtokenか検証
		str, err := getTokenString(token)
		if err != nil {
			presenter.BadRequestError(w, err)
			return 
		}

		// tokenからclaimsに変換
		claims, err := value.Parse(str)
		if err != nil {
			presenter.BadRequestError(w, err)
			return
		}

		inp := user.FindNameInput{
			ID: claims.User.ID,
		}
		_, aerr := m.usecase.FindName(inp)
		if aerr != nil {
			presenter.ApplicationException(w, aerr)
			return
		}

		// ctxに"key=id, value=claims.User.ID"をセットする
		ctx = context.WithValue(ctx, "id", claims.User.ID)
		// 次のhandlerにrequestにctxを持たせる
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func getTokenString(token string) (string, error) {
	if len(token) < 8 {
		return "", errors.New("Invalid Authorization header")
	}
	if token[:7] != "Bearer "  {
		return "", errors.New("Authorization Header is not 'Bearer' token")
	}
	return token[7:], nil
}