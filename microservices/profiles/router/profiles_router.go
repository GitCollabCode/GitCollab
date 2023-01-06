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
	r.With(p.MiddleWareValidateProfile).Post("/", p.PostProfile)

	r.Get("/search", p.SearchProfile)
	r.Route("/{username}", func(r chi.Router) {
		r.Get("/", p.GetProfile)
		r.Delete("/", p.DeleteProfile)
	})

	r.Route("/skills", func(r chi.Router) {
		r.Use(jwt.JWTBlackList(p.Pd.PDriver))
		r.Use(jwtConf.VerifyJWT(p.Pd.PDriver.Log))
		r.Patch("/", p.PatchSkills)
		r.Delete("/", p.DeleteSkills)
	})
	return r
}
