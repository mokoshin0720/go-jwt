package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"github.com/ispec-inc/sample/pkg/domain/model"
	"github.com/ispec-inc/sample/pkg/presenter"
	"github.com/ispec-inc/sample/pkg/registry"
	"github.com/ispec-inc/sample/pkg/view"
	"github.com/ispec-inc/sample/src/user"
)

type handler struct {
	usecase user.Usecase
}

func NewHandler(repo registry.Repository) handler {
	usecase := user.NewUsecase(repo)
	return handler{usecase}
}

func (h handler) GetName(w http.ResponseWriter, r * http.Request) {
	// URLのパラメータからidを抽出
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		presenter.BadRequestError(w, err)
		return
	}

	// usecaseで記載したinput.go
	inp := user.FindNameInput{
		ID: int64(id),
	}
	out, aerr := h.usecase.FindName(inp)
	// daoで記載したエラ-をhttp.ResponseWriterで記載
	if aerr != nil {
		presenter.ApplicationException(w, aerr)
		return 
	}

	ures := view.NewUserName(out.User)
	res := GetNameResponse{
		UserName: ures,
	}
	presenter.Response(w, res)
}

func (h handler) AddName(w http.ResponseWriter, r *http.Request) {
	var request addNameRequest
	// jsonに変換できるかどうか
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		presenter.BadRequestError(w, err)
		return 
	}
	// validationに関するエラー
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		presenter.BadRequestError(w, err)
		return
	}

	u := model.User{
		Name: request.Name,
		Email: request.Email,
		Password: request.Password,
	}
	inp := user.AddNameInput{
		User: u,
	}
	out, aerr := h.usecase.AddName(inp)
	if aerr != nil {
		presenter.ApplicationException(w, aerr)
		return 
	}

	ures := view.NewUserName(out.User)
	res := AddNameResponse{
		UserName: ures,
	}

	presenter.Response(w, res)
}

func (h handler) GetJwt(w http.ResponseWriter, r *http.Request) {
	var request jwtRequest
	// jsonに関するエラーハンドリング
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		presenter.BadRequestError(w, err)
		return
	}
	// requestのvalidateに関するエラーハンドリング
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		presenter.BadRequestError(w, err)
		return
	}

	inp := user.JwtLoginInput{
		Email: request.Email,
		Password: request.Password,
	}

	out, aerr := h.usecase.JwtLogin(inp)
	if aerr != nil {
		presenter.ApplicationException(w, aerr)
		return 
	}

	res := GetJWTResponse{
		Token: out.Token,
	}

	presenter.Response(w, res)
}

func (h handler) Restricted(w http.ResponseWriter, r *http.Request) {
	var request restrictedRequest
	// jsonに変換する際のエラーハンドリング
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		presenter.BadRequestError(w, err)
		return
	}
	// validationのエラーハンドリング
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		presenter.BadRequestError(w, err)
		return
	}
	
	inp := user.LoginUserInput{
		TokenString: request.Token,
	}
	out, aerr := h.usecase.LoginUser(inp)
	if aerr != nil {
		presenter.ApplicationException(w, aerr)
		return 
	}

	ures := view.NewUserName(out.User)
	res := GetNameResponse{
		UserName: ures,
	}
	presenter.Response(w, res)
}