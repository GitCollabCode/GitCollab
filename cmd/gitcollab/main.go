// Test main.go file to create docker image
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/GitCollabCode/GitCollab/internal/db"
	authHandlers "github.com/GitCollabCode/GitCollab/microservices/authentication/handlers"
	authRouter "github.com/GitCollabCode/GitCollab/microservices/authentication/router"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
)

func verifyEnv(env ...string) error {
	for _, str := range env {
		if str == "" {
			return fmt.Errorf("failed to import an environment variable")
		}
	}
	return nil
}

func main() {
	r := chi.NewRouter()

	// initialize logger
	log := logrus.New()
	log.Info("Starting Logger!")

	// get environment variables
	clientID := os.Getenv("GITHUB_CLIENTID")
	gitRedirect := os.Getenv("REACT_APP_REDIRECT_URI")

	// check environment variables
	if err := verifyEnv(clientID, gitRedirect); err != nil {
		log.Panic(err.Error())
		return
	}

	// register middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.StripSlashes)
	r.Use(cors.Handler(cors.Options{
		//AllowedOrigins:   []string{"https://localhost"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// create db drivers
	authDB, err := db.ConnectPostgres(os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Error(err)
		return
	}
	defer authDB.Connection.Close(context.Background())

	// Set routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi from Git Collab"))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("pong!"))
	})

	// register all sub routers
	auth := authHandlers.NewAuth(log, authDB, clientID, gitRedirect)
	authRouter.InitAuthRouter(r, auth)

	// Start server
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = ":8080"
	}

	http.ListenAndServe(httpPort, r)
}
