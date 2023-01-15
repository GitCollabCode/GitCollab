package router

import (
	"time"

	"github.com/GitCollabCode/GitCollab/internal/gitauth"
	jwt "github.com/GitCollabCode/GitCollab/internal/jwt"
	profileData "github.com/GitCollabCode/GitCollab/microservices/profiles/data"
	"github.com/GitCollabCode/GitCollab/microservices/projects/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
)

const MAX_GIT_TRANSACTIONS = 5 // per min

// Defines Project service endpoints
func ProjectRouter(project *handlers.Projects, profiles *profileData.ProfileData) chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(jwt.JWTBlackList(project.PgConn))
		r.Use(project.JwtConf.VerifyJWT(project.Log))
		r.Use(gitauth.GitClient(profiles))

		r.Route("/github", func(r chi.Router) {
			r.Use(httprate.LimitByIP(MAX_GIT_TRANSACTIONS, 1*time.Minute))
			r.Get("/user-repos", project.GetUserRepos)
			r.Get("/repo-info", project.GetRepoInfo)
			r.Get("/repo-issues", project.GetRepoIssues)
		})

		r.Post("/create-project", project.CreateProject)
	})
	r.Get("/user-projects", project.GetUserProjects)
	r.Get("/projects-issues", project.GetProjectIssues)

	return r
}
