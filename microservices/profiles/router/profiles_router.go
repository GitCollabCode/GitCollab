package router

import (
	"github.com/GitCollabCode/GitCollab/internal/jwt"
	genericMiddleware "github.com/GitCollabCode/GitCollab/internal/middleware"
	"github.com/GitCollabCode/GitCollab/microservices/profiles/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Defines Profiles service endpoints
func ProfileRouter(p *handlers.Profiles, jwtConf *jwt.GitCollabJwtConf) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentEncoding("application/json"))
	r.Use(genericMiddleware.SetContentType("application/json"))

	// swagger:route POST /profiles Profiles createProfileRequest
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
	r.Post("/", p.PostProfile)

	// swagger:route POST /profiles/search Profiles profileSearchRequest
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
	r.Post("/search", p.SearchProfile)

	// swagger:route Get /profiles/get-skills Profiles profileSkillsRequest
	//
	// Get available skills.
	//
	// Get a list of available skills.
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: searchProfilesResponse
	//		 204: description:No Profiles found
	r.Get("/get-skills", p.GetSkillList)

	// swagger:route Get /profile/get-languages Profiles profileLanguageRequest
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
	r.Get("/get-languages", p.GetLanguageList)

	r.Route("/{username}", func(r chi.Router) {
		// swagger:route GEzusername} Profiles getProfile
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
