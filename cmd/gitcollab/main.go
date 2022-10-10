// Test main.go file to create docker image
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/GitCollabCode/GitCollab/internal/jwt"
	authHandlers "github.com/GitCollabCode/GitCollab/microservices/authentication/handlers"
	authRouter "github.com/GitCollabCode/GitCollab/microservices/authentication/router"
	profilesHandlers "github.com/GitCollabCode/GitCollab/microservices/profiles/handlers"
	profilesRouter "github.com/GitCollabCode/GitCollab/microservices/profiles/router"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
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
	logger := logrus.New()
	logger.Out = os.Stdout

	logger.Info("Starting Logger!")

	// get environment variables
	clientID := os.Getenv("GITHUB_CLIENTID")
	clientSecret := os.Getenv("GITHUB_SECRET")
	gitCollabSecret := os.Getenv("GITCOLLAB_SECRET")
	gitRedirect := os.Getenv("REACT_APP_REDIRECT_URI")

	// check environment variables
	if err := verifyEnv(clientID, gitRedirect); err != nil {
		logger.Panic(err.Error())
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

	// create oauth config for github
	var GitOauthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,                        // maybe store?
		Scopes:       []string{"user:email", "user:name"}, // verify what we need
		Endpoint:     githuboauth.Endpoint,
	}

	// create gitcollab jwt conf
	jwtConf := jwt.NewGitCollabJwtConf(gitCollabSecret)

	// create db drivers
	dbDriver, err := db.ConnectPostgres(os.Getenv("POSTGRES_URL"))
	if err != nil {
		logger.Error(err)
		return
	}
	defer dbDriver.Connection.Close(context.Background())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi from Git Collab"))
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("pong!"))
	})

	// register all sub routers
	auth := authHandlers.NewAuth(dbDriver, logger, GitOauthConfig, gitRedirect)
	authRouter.InitAuthRouter(r, auth, jwtConf)

	profiles := profilesHandlers.NewProfiles(logger)
	profilesRouter.InitRouter(r, profiles)

	// Start server
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = ":8080"
	}

	//TODO: add debug flag condition for this
	r.Mount("/debug", middleware.Profiler())

	s := http.Server{
		Addr:         httpPort,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
		//ErrorLog:   log,
	}

	go func() {
		logger.Infof("Starting server on port %s", httpPort)

		err := s.ListenAndServe()
		if err != nil {
			logger.Error("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// Trap interupt signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received
	sig := <-c
	logger.Info("Got interrupt signal:", sig)

	// Gracefully server shutdown wait 30 seconds for any ongoing operations to complete
	// Check if the docker container is killed in away that allows for this to happen
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
