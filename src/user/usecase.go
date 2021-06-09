package user

import (
	"fmt"

	"github.com/ispec-inc/sample/pkg/apperror"
	"github.com/ispec-inc/sample/pkg/domain/repository"
	"github.com/ispec-inc/sample/pkg/password"
	"github.com/ispec-inc/sample/pkg/registry"
	"github.com/ispec-inc/sample/pkg/value"
)

type Usecase struct {
	user repository.User // 全てのメソッドを記載したinterface
}

func NewUsecase(repo registry.Repository) Usecase {
	return Usecase{
		user: repo.NewUser(), // daoの初期化？
	}
}

func (use Usecase) FindName(inp FindNameInput) (out FindNameOutput, aerr apperror.Error) {
	u, aerr := use.user.Find(inp.ID)
	if aerr != nil {
		return 
	}
	out.User = u
	return out, nil
}

func (use Usecase) AddName(inp AddNameInput) (out AddNameOutput, aerr apperror.Error) {
	// パスワードのハッシュ化
	pwd, err := password.Encrypt(inp.User.Password)
	if err != nil {
		aerr = apperror.New(apperror.CodeError, err)
		return 
	}
	inp.User.Password = pwd

	aerr = use.user.Create(inp.User)
	if aerr != nil {
		return
	}

	// 実際にUserが作られたかどうか
	u, aerr := use.user.FindByEmail(inp.User.Email)
	if aerr != nil {
		return
	}
	out.User = u
	
	return out, nil
}

// EmailとPasswordからString(token)を生成
func (use Usecase) JwtLogin(inp JwtLoginInput) (out JwtLoginOutput, aerr apperror.Error) {
	// 指定したEmailをもつUserが存在するかどうか
	u, aerr := use.user.FindByEmail(inp.Email)
	if aerr != nil {
		return 
	}
	
	// ハッシュ化されたDB上のパスワードと、inp.Passwordをハッシュ化したものを比較
	ok := password.Compare(u.Password, inp.Password)
	if !ok {
		aerr = apperror.New(apperror.CodeInvalid, fmt.Errorf("パスワードが正しくありません。"))
		return
	}

	claims := value.NewClaims(u)
	token, err := value.NewToken(claims)
	if err != nil {
		aerr = apperror.New(apperror.CodeError, err)
		return
	}

	out.Token = token.String()
	
	return out, nil
}

func (use Usecase) LoginUser(inp LoginUserInput)(out LoginUserOutput, aerr apperror.Error) {
	// Tokenを渡すとユーザーのEmailを返してくれる
	claims, err := value.Parse(inp.TokenString)
	if err != nil {
		aerr = apperror.New(apperror.CodeError, err)
		return
	}

	// tokenユーザーがいるかどうか
	u, aerr := use.user.Find(claims.User.ID)
	if aerr != nil {
		return 
	}

	out.User = u

	return out, nil
}