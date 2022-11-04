package router

import (
	"github.com/GitCollabCode/GitCollab/internal/jwt"
	"github.com/GitCollabCode/GitCollab/microservices/authentication/handlers"
	"github.com/go-chi/chi/v5"
)

func InitAuthRouter(r chi.Router, auth *handlers.Auth, conf *jwt.GitCollabJwtConf) {
	// add all routes and handlers below

	r.Route("/project", func(r chi.Router) {
		r.Get("/get", )
		r.Post("/signin", auth.LoginHandler)
		r.Get("/logout", auth.LogoutHandler)
	})

}