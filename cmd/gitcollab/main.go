// GitCollab API
//
// GitCollab API Swagger documentation.
//
// Terms of Service:
//
// There is currently no Terms of Service.
//
//	Schemes: http
//	BasePath: /api
//	Version: 1.0.0
//	License: MIT http://opensource.org/licenses/MIT
//	Host: gitcollab.tridentshark.com
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Security:
//	- bearer
//
//	SecurityDefinitions:
//	bearer:
//	     type: apiKey
//	     name: Authorization
//	     in: header
//
// swagger:meta
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/GitCollabCode/GitCollab/internal/jwt"
	authHandlers "github.com/GitCollabCode/GitCollab/microservices/authentication/handlers"
	authRouter "github.com/GitCollabCode/GitCollab/microservices/authentication/router"
	"github.com/GitCollabCode/GitCollab/microservices/profiles/data"
	profilesHandlers "github.com/GitCollabCode/GitCollab/microservices/profiles/handlers"
	profilesRouter "github.com/GitCollabCode/GitCollab/microservices/profiles/router"
	projectData "github.com/GitCollabCode/GitCollab/microservices/projects/data"
	project "github.com/GitCollabCode/GitCollab/microservices/projects/handlers"
	projectsRouter "github.com/GitCollabCode/GitCollab/microservices/projects/router"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	logger.SetReportCaller(true)

	logger.Info("Starting Logger!")

	// get environment variables
	clientID := os.Getenv("GITHUB_CLIENTID")
	clientSecret := os.Getenv("GITHUB_SECRET")
	gitCollabSecret := os.Getenv("GITCOLLAB_SECRET")
	gitRedirect := os.Getenv("REACT_APP_REDIRECT_URI")
	dbUrl := os.Getenv("POSTGRES_URL")
	httpPort := os.Getenv("HTTP_PORT")

	// check environment variables
	if err := verifyEnv(clientID, gitRedirect, dbUrl); err != nil {
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
		AllowedMethods:   []string{"GET", "PATCH", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// create oauth config for github
	var GitOauthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"user:email", "user:name"},
		Endpoint:     githuboauth.Endpoint,
	}

	// create gitcollab jwt conf, for midddleware
	jwtConf := jwt.NewGitCollabJwtConf(gitCollabSecret)

	// create db drivers
	dbDriver, err := db.NewPostgresDriver(dbUrl, logger)
	if err != nil {
		logger.Error(err)
		return
	}
	defer dbDriver.Pool.Close()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("hi from Git Collab"))
		if err != nil {
			log.Panic(err)
		}
	})

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		_, err := w.Write([]byte("pong!"))
		if err != nil {
			log.Panic(err)
		}
	})

	// register all sub routers under /api
	r.Route("/api", func(r chi.Router) {
		// authentication subrouter
		auth := authHandlers.NewAuth(dbDriver, logger, GitOauthConfig, gitRedirect, gitCollabSecret)
		r.Mount("/auth", authRouter.AuthRouter(auth))

		// profiles subrouter
		pd := data.NewProfileData(dbDriver)
		profiles := profilesHandlers.NewProfiles(logger, pd)
		r.Mount("/profile", profilesRouter.ProfileRouter(profiles, jwtConf))

		projectD := projectData.NewProjectData(dbDriver)
		p := project.NewProjects(dbDriver, projectD, jwtConf, logger)
		projectD.AddProject(1234, "test123", "asdqrtqf qerw12 123 asd")
		r.Mount("/project", projectsRouter.ProjectRouter(p, pd))

		// test routes
		r.Route("/test", func(r chi.Router) {
			r.Use(jwt.JWTBlackList(dbDriver))
			r.Get("/test-blacklist", func(w http.ResponseWriter, r *http.Request) {
				_, err := w.Write([]byte("cheese"))
				if err != nil {
					logger.Info("Burgre King")
				}
			})
		})
	})

	// Start server
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
			logger.Errorf("Error starting server: %s\n", err)
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
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = s.Shutdown(ctx)
	if err != nil {
		log.Panic(err)
	}
}
