package main

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/ispec-inc/sample/cmd/api/adapter/invitation"
	"github.com/ispec-inc/sample/cmd/api/adapter/user"
	"github.com/ispec-inc/sample/cmd/api/middleware"
	"github.com/ispec-inc/sample/pkg/presenter"
	"github.com/ispec-inc/sample/pkg/registry"
)

func NewRouter(repo registry.Repository) http.Handler {
	r := chi.NewRouter()

	// トークン検証をするmiddleware
	authmiddleware := middleware.NewAuth(repo)
	invitationHandler := invitation.NewHandler(repo)
	userHandler := user.NewHandler(repo)

	r = commonMiddleware(r)

	r.Mount("/invitations", invitation.NewRouter(invitationHandler))
	// エンドポイントusersにrequestを投げるときは、authmiddlewareを通る
	r.Mount("/users", user.NewRouter(userHandler, authmiddleware))
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		presenter.Success(w)
	})

	return r
}
