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
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		presenter.BadRequestError(w, err)
		return
	}

	inp := user.FindNameInput{
		ID: int64(id),
	}
	out, aerr := h.usecase.FindName(inp)
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
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		presenter.BadRequestError(w, err)
		return 
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		presenter.BadRequestError(w, err)
		return 
	}

	u := model.User{
		ID: request.ID,
		Name: request.Name,
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