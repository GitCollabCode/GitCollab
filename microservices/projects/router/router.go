package router

import (
	"time"

	"github.com/GitCollabCode/GitCollab/microservices/projects/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/httprate"
)

const MAX_GIT_TRANSACTIONS = 5 // per min

func ProjectRouter(project *handlers.Projects) chi.Router {
	r := chi.NewRouter()
	// github related endpoints, should not be called often,
	r.Route("/github", func(r chi.Router) {
		r.Use(httprate.LimitByIP(MAX_GIT_TRANSACTIONS, 1*time.Minute))
		r.Get("/user-repos", project.GetUserRepos)
		r.Get("/repo-info", project.GetRepoInfo)
		r.Get("/repo-issues", project.GetRepoIssues)
	})

	// our endpoints
	r.Post("/create-project", project.CreateProject)
	r.Patch("/project-description", project.PatchProjectDescription)
	r.Get("/user-projects", project.GetUserProjects)
	r.Get("/projects-issues", project.GetProjectIssues)

	return r
}
