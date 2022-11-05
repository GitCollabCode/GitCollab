package router

import (
	"fmt"
	"net/http"
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

func TestRouter(auth *handlers.Auth) chi.Router {
	r := chi.NewRouter()
	r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("I got here")
		auth.GithubRedirectHandler(w, r)
	})
	return r
}

func AuthRouter(auth *handlers.Auth) chi.Router {
	r := chi.NewRouter()
	r.Use(httprate.LimitByIP(MAX_SIGNIN_MIN, 1*time.Minute))
	r.Get("/redirect-url", auth.GithubRedirectHandler)
	r.Post("/signin", auth.LoginHandler)
	r.Get("/logout", auth.LogoutHandler)
	return r
}
