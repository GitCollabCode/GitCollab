package router

import (
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/jwt"
	"github.com/GitCollabCode/GitCollab/microservices/authentication/handlers"
	"github.com/go-chi/chi/v5"
)

func InitAuthRouter(r chi.Router, auth *handlers.Auth) {
	// add all routes and handlers below
	r.Route("/auth", func(r chi.Router) {
		r.Get("/redirect-url", auth.GithubRedirectHandler)
		r.Post("/signin", auth.LoginHandler)
		r.Get("/logout", auth.LogoutHandler)
	})

	r.Route("/poop", func(r chi.Router) {
		r.Use(jwt.JWTBlackList(auth.PgConn, auth.Log))
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("I LOVE RACISM"))
		})
	})
}
