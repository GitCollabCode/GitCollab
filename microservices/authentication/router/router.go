package router

import (
	"time"

	"github.com/GitCollabCode/GitCollab/microservices/authentication/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
)

const MAX_SIGNIN_MIN = 10

func AuthRouter(auth *handlers.Auth) chi.Router {
	r := chi.NewRouter()
	r.Use(httprate.LimitByIP(MAX_SIGNIN_MIN, 1*time.Minute))
	r.Get("/redirect-url", auth.GithubRedirectHandler)
	r.Post("/signin", auth.LoginHandler)
	r.Post("/logout", auth.LogoutHandler)
	return r
}
