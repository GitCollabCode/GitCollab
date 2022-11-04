package router

import (
	"time"

	"github.com/GitCollabCode/GitCollab/internal/jwt"
	"github.com/GitCollabCode/GitCollab/microservices/authentication/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
)

const MAX_SIGNIN_MIN = 10

func InitAuthRouter(r chi.Router, auth *handlers.Auth, conf *jwt.GitCollabJwtConf) {
	// add all routes and handlers below

	r.Route("/auth", func(r chi.Router) {
		r.Get("/redirect-url", auth.GithubRedirectHandler)
		r.Route("/signin", func(r chi.Router) {
			// limit singins per minute
			r.Use(httprate.LimitByIP(MAX_SIGNIN_MIN, 1*time.Minute))
			r.Post("/", auth.LoginHandler)
		})
		r.Get("/logout", auth.LogoutHandler)
	})

}
