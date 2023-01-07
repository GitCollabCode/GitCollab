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
	r.With(p.MiddleWareValidateProfile).Post("/", p.PostProfile)

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

	return r
}
