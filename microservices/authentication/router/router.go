package router

import (
	"fmt"
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/jwt"
	"github.com/GitCollabCode/GitCollab/microservices/authentication/handlers"
	"github.com/go-chi/chi/v5"
)

func InitAuthRouter(r chi.Router, auth *handlers.Auth, conf *jwt.GitCollabJwtConf) {
	// add all routes and handlers below

	r.Route("/auth", func(r chi.Router) {
		//
		r.Get("/redirect-url", auth.GithubRedirectHandler)
		r.Post("/signin", auth.LoginHandler)
		r.Get("/logout", auth.LogoutHandler)
		r.Route("/test", func(r chi.Router) {
			r.Use(conf.VerifyJWT(auth.Log))
			fmt.Println("Taco bell crunch wrap supreme")
			r.Get("/verify", func(w http.ResponseWriter, r *http.Request) {
				//w.Write([]byte("crungchy chickem"))
				auth.TestHandler(w, r)
			})
		})
	})

}
