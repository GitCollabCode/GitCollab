// Test main.go file to create docker image
package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/GitCollabCode/GitCollab/internal/jwt"
	"github.com/GitCollabCode/GitCollab/microservices/authentication/router"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	// initialize logger
	log := logrus.New()
	log.Info("Starting Logger!")

	// create db drivers
	authDB, err := db.ConnectPostgres(os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Error(err)
		return
	} 
	defer authDB.Connection.Close(context.Background())

	// add middleware for JWT
	SECRET := os.Getenv("JWT_SECRET")
	tokenAuth := jwtauth.New("HS256", []byte(SECRET), nil)

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	// middleware for blacklist
	r.Use(jwt.JWTBlackList(tokenAuth, authDB, log))

	// midleware for jwt
	r.Use(jwtauth.Verifier(tokenAuth))

	// init authentication microservice
	router.InitAuth(tokenAuth, log)	

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi from Git Collab"))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("pong!"))
	})

	r.Mount("/auth", router.AuthRouter())

	// Start server
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = ":8080"
	}
	http.ListenAndServe(httpPort, r)
}
