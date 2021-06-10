package user

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ispec-inc/sample/cmd/api/middleware"
)

func NewRouter(handler handler, middleware middleware.Auth) http.Handler {
	r := chi.NewRouter()

	r.Post("/", handler.AddName)
	r.Post("/jwt", handler.GetJwt)
	r.Post("/private", handler.Restricted)

	// トークン検証を挟む
	r.Group(func(r chi.Router) {
		r.Use(middleware.VerifyToken)
		r.Get("/{id}", handler.GetName)
	})
	return r
}