package router

import (
	"github.com/GitCollabCode/GitCollab/internal/jwt"
	"github.com/GitCollabCode/GitCollab/microservices/authentication/handlers"
	"github.com/go-chi/chi/v5"
)

func InitAuthRouter(r chi.Router, projects *handlers.Projects, conf *jwt.GitCollabJwtConf) {
	// add all routes and handlers below

	r.Route("/project", func(r chi.Router) {
		r.Get("/project", projects.ProjectsForProfileHandler )
		
	})

}