package router

import (
	"github.com/GitCollabCode/GitCollab/internal/jwt"
	genericMiddleware "github.com/GitCollabCode/GitCollab/internal/middleware"
	"github.com/GitCollabCode/GitCollab/microservices/tasks/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Defines Profiles service endpoints
func TasksRouter(t *handlers.Tasks, jwtConf *jwt.GitCollabJwtConf) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentEncoding("application/json"))
	r.Use(genericMiddleware.SetContentType("application/json"))

	// swagger:route POST /projects/{project-name}/tasks/new Profiles createProfileRequest
	//
	// Create a user profile.
	//
	// Takes in user information to create a new profile, only used for testing.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: messageResponse
	r.Post("/", t.CreateTask)

	// swagger:route POST /projects/{project-name}/tasks Profiles profileSearchRequest
	//
	// Search registered profiles.
	//
	// Get profile information based on input search parameters.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: searchProfilesResponse
	//		 204: description:No Profiles found
	r.Post("/search", t.GetTasks)

	// swagger:route Get /project/{project-name}/tasks/ Profiles profileLanguageRequest
	//
	// Get available languages.
	//
	// Get a list of available languages
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: messageResponse
	r.Get("/get-languages", t.DeleteTask)

	r.Route("/{username}", func(r chi.Router) {
		// swagger:route GET /profiles/{username} Profiles getProfile
		//
		// Get GitCollab profile.
		//
		//     Parameters:
		//       + name: username
		//         in: path
		//         description: Target user
		//         required: true
		//         type: string
		//
		//     Produces:
		//     - application/json
		//
		//     Responses:
		//       200: getProfileResponse
		//		 404: description:Profile not found
		r.Get("/", p.GetProfile)

		r.Group(func(r chi.Router) {
			// setup private routes
			r.Use(jwt.JWTBlackList(p.Pd.PDriver))
			r.Use(jwtConf.VerifyJWT(p.Pd.PDriver.Log))

			// swagger:route DELETE /profiles/{username} Profiles deleteProfile
			//
			// Delete GitCollab user.
			//
			// Removes registered GitCollab user from db, only used for testing, should not be exposed.
			//
			//     Parameters:
			//       + name: username
			//         in: path
			//         description: Target user
			//         required: true
			//         type: string
			//
			//     Produces:
			//     - application/json
			//
			//     Responses:
			//       200: messageResponse
			//		 404: description:Profile not found
			r.Delete("/", p.DeleteProfile)
		})
	})

	r.Route("/skills", func(r chi.Router) {
		r.Use(jwt.JWTBlackList(p.Pd.PDriver))
		r.Use(jwtConf.VerifyJWT(p.Pd.PDriver.Log))

		// swagger:route PATCH /profiles/skills Profiles profileSkillsRequest
		//
		// Patch profile skills.
		//
		// Append provided skills to the callers profile.
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
		r.Patch("/", p.PatchSkills)

		// swagger:route POST /profiles/skills Profiles profileSkillsRequest
		//
		// Delete profile skills.
		//
		// Delete provided skills from the callers profile.
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
		r.Delete("/", p.DeleteSkills)
	})

	r.Route("/languages", func(r chi.Router) {
		r.Use(jwt.JWTBlackList(p.Pd.PDriver))
		r.Use(jwtConf.VerifyJWT(p.Pd.PDriver.Log))

		// swagger:route PATCH /profile/languages Profiles profileSkillsRequest
		//
		// Patch profile Languages.
		//
		// Append provided Languages to the callers profile.
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
		r.Patch("/", p.PatchLanguages)

		// swagger:route POST  /profile/languages Profiles profileSkillsRequest
		//
		// Delete profile languages.
		//
		// Delete provided skills from the callers profile.
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
		r.Delete("/", p.DeleteLanguages)
	})

	return r
}
