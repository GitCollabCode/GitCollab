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

// ProjectRouter serve projects api endpoints
func ProjectRouter(p *handlers.Projects, profiles *profileData.ProfileData, jwtConf *jwt.GitCollabJwtConf) chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(jwt.JWTBlackList(p.PgConn))
		r.Use(jwtConf.VerifyJWT(p.Log))
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
			r.Get("/user-repos", p.GetUserRepos)

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
			r.Get("/repo-info", p.GetRepoInfo)

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
			r.Get("/repo-issues", p.GetRepoIssues)
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
		r.Post("/create-project", p.CreateProject)
	})

	r.Route("/{project-name}", func(r chi.Router) {

		r.Route("/tasks", func(r chi.Router) {
			r.Get("/", p.GetTasks)

			r.Route("/{task-id}", func(r chi.Router) {
				r.Get("/", p.GetTask)

				r.Group(func(r chi.Router) {
					r.Use(jwt.JWTBlackList(p.PgConn))
					r.Use(jwtConf.VerifyJWT(p.Log))
					r.Delete("/", p.DeleteTask)
					r.Patch("/", p.EditTask)
				})
			})

			r.Group(func(r chi.Router) {
				r.Use(jwt.JWTBlackList(p.PgConn))
				r.Use(jwtConf.VerifyJWT(p.Log))
				r.Post("/new", p.CreateTask)
			})
		})

		r.Get("/", p.GetProject)

	})

	// swagger:route Get /projects/user-projects Projects userProjectsReq
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
	r.Get("/user-projects", p.GetUserProjects)

	// swagger:route get /projects/search-projects Projects
	//
	// Get GitCollab projects, just random ones for now
	//
	// Retrieve list of GitCollab projects
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: messageResponse
	r.Get("/search-projects", p.GetSearchProjects)
	return r
}
