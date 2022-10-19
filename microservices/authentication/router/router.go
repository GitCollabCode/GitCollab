package router

import (
	"github.com/GitCollabCode/GitCollab/internal/jwt"
	"github.com/GitCollabCode/GitCollab/microservices/authentication/handlers"
	"github.com/go-chi/chi/v5"
)

func InitAuthRouter(r chi.Router, auth *handlers.Auth, conf *jwt.GitCollabJwtConf) {
	// add all routes and handlers below

	r.Route("/auth", func(r chi.Router) {
		r.Get("/redirect-url", auth.GithubRedirectHandler)
		r.Post("/signin", auth.LoginHandler)
		r.Get("/logout", auth.LogoutHandler)
	})

}
