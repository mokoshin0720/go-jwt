package user

import (
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter(handler handler) http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", handler.GetName)
	r.Post("/", handler.AddName)
	r.Get("/{email:[a-z-]+}", handler.GetPassword)
	r.Post("/jwt", handler.GetJwt)
	r.Post("/private", handler.Restricted)
	return r
}