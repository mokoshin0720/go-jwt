package main

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/ispec-inc/sample/cmd/api/adapter/invitation"
	"github.com/ispec-inc/sample/cmd/api/adapter/user"
	"github.com/ispec-inc/sample/pkg/presenter"
	"github.com/ispec-inc/sample/pkg/registry"
)

func NewRouter(repo registry.Repository) http.Handler {
	r := chi.NewRouter()

	invitationHandler := invitation.NewHandler(repo)
	userHandler := user.NewHandler(repo)

	r = commonMiddleware(r)

	r.Mount("/invitations", invitation.NewRouter(invitationHandler))
	r.Mount("/users", user.NewRouter(userHandler))
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		presenter.Success(w)
	})

	return r
}
