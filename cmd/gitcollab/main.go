// Test main.go file to create docker image
package main

import (
	"net/http"
	"os"
	"time"

	"github.com/GitCollabCode/GitCollab/microservices/authentication/router"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	// add middleware for JWT
	SECRET := os.Getenv("JWT_SECRET")
	tokenAuth := jwtauth.New("HS256", []byte(SECRET), nil)
	r.Use(jwtauth.Verifier(tokenAuth))

	// add middleware fro JWT Blacklist

	// initialize logger
	log := logrus.New()
	log.Info("Starting Logger!")

	// get port for backend
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = ":8080"
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi from Git Collab"))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("pong!"))
	})

	r.Mount("/auth", router.AuthRouter())
	http.ListenAndServe(httpPort, r)
}
