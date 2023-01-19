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

// ProjectRouter defines Project service endpoints
func ProjectRouter(project *handlers.Projects, profiles *profileData.ProfileData, jwtConf *jwt.GitCollabJwtConf) chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(jwt.JWTBlackList(project.PgConn))
		r.Use(jwtConf.VerifyJWT(project.Log))
		r.Use(gitauth.GitClient(profiles))

		r.Route("/github", func(r chi.Router) {
			r.Use(httprate.LimitByIP(MAX_GIT_TRANSACTIONS, time.Minute))

			// swagger:route GET /projects/github/user-repos Projects GitHub githubGetUserRepos
			//
			// Get GitHub repos owned by user.
			//
			// Retrieve a list of GitHub repos owned by a user.
			//
			//     Parameters:
			//       + name: Authorization
			//         in: header
			//         description: User JWT
			//         required: true
			//         type: string
			//
			//     Produces:
			//     - application/json
			//
			//     Responses:
			//       200: reposGetResp
			r.Get("/user-repos", project.GetUserRepos)

			// swagger:route GET /projects/github/repo-info Projects GitHub repoInfoReq
			//
			// Get repo info from GitHub.
			//
			// Retrieve detailed information about a select repo from GitHub.
			//
			//     Parameters:
			//       + name: Authorization
			//         in: header
			//         description: User JWT
			//         required: true
			//         type: string
			//
			//     Consumes:
			//     - application/json
			//
			//     Produces:
			//     - application/json
			//
			//     Responses:
			//       200: repoInfoResp
			r.Get("/repo-info", project.GetRepoInfo)

			// swagger:route GET /projects/github/repo-issues Projects GitHub repoInfoReq
			//
			// Get issues under a repo from GitHub.
			//
			// Retrieve a list of all open issues under a repo from GitHub.
			//
			//     Parameters:
			//       + name: Authorization
			//         in: header
			//         description: User JWT
			//         required: true
			//         type: string
			//
			//     Consumes:
			//     - application/json
			//
			//     Produces:
			//     - application/json
			//
			//     Responses:
			//       200: repoIssueResp
			r.Get("/repo-issues", project.GetRepoIssues)
		})

		// swagger:route POST /projects/create-project Projects repoInfoReq
		//
		// Create a GitCollab project.
		//
		// Create a GitCollab project based on a select repo.
		//
		//     Parameters:
		//       + name: Authorization
		//         in: header
		//         description: User JWT
		//         required: true
		//         type: string
		//
		//     Consumes:
		//     - application/json
		//
		//     Produces:
		//     - application/json
		//
		//     Responses:
		//       200: messageResponse
		r.Post("/create-project", project.CreateProject)
	})

	// swagger:route POST /projects/user-projects Projects userProjectsReq
	//
	// Get GitCollab projects owned by a user.
	//
	// Retrieve list of GitCollab projects associated to a given user.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: messageResponse
	r.Get("/user-projects", project.GetUserProjects)

	return r
}
