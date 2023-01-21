package router

import (
	"time"

	"github.com/GitCollabCode/GitCollab/microservices/authentication/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httprate"
)

const MAX_SIGNIN_MIN = 10

// AuthRouter serve authentication api endpoints
func AuthRouter(auth *handlers.Auth) chi.Router {
	r := chi.NewRouter()
	r.Use(httprate.LimitByIP(MAX_SIGNIN_MIN, 1*time.Minute))

	// swagger:route GET /auth/redirect-url Auth getRedirectUrl
	//
	// Get GitCollab redirect url.
	//
	// Get the redirect url for GitHub that corresponds to the GitCollab app GitHub registration.
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: redirectResponse
	r.Get("/redirect-url", auth.GithubRedirectHandler)

	// swagger:route POST /auth/signin Auth githubOAuthRequest
	//
	// User Login through GitHub.
	//
	// Fetch user GitHub access token using users GitHub OAuth code and returns JWT.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Responses:
	//       200: loginResponse
	//		 400: description:Missing GitHub OAuth code
	//       401: description:Invalid GitHub OAuth code
	r.Post("/signin", auth.LoginHandler)

	// swagger:route POST /auth/logout Auth logout
	//
	// User logout.
	//
	// Logs user out and invalidates their currently valid JWT.
	//
	//     Parameters:
	//       + name: Authorization
	//         in: header
	//         description: User JWT
	//         required: true
	//         type: string
	//
	//     Responses:
	//       200: description:Successful logout
	r.Post("/logout", auth.LogoutHandler)
	return r
}
