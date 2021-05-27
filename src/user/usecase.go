package user

import (
	"github.com/ispec-inc/sample/pkg/domain/repository"
	"github.com/ispec-inc/sample/pkg/registry"
	"github.com/ispec-inc/sample/pkg/apperror"
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
	aerr = use.user.Create(inp.User)
	if aerr != nil {
		return 
	}

	u, aerr := use.user.Find(inp.User.ID)
	if aerr != nil {
		return 
	}
	out.User = u
	
	return out, nil
}

func (use Usecase) FindPassword(inp FindPasswordInput) (out FindPasswordOutput, aerr apperror.Error) {
	u, aerr := use.user.EmailPass(inp.Email)
	if aerr != nil {
		return 
	}
	out.User = u
	return out, nil
}

func (use Usecase) JwtLogin(inp JwtLoginInput) (out JwtLoginOutput, aerr apperror.Error) {
	u, aerr := use.user.Publish(inp.Email, inp.Password)
	if aerr != nil {
		return 
	}
	out.User = u
	
	return out, nil
}

func (use Usecase) LoginUser(inp LoginUserInput)(out LoginUserOutput, aerr apperror.Error) {
	u, aerr := use.user.ParseToken(inp.TokenString)
	if aerr != nil {
		return 
	}
	out.User = u
	return out, nil
}