package router

import (
	"github.com/GitCollabCode/GitCollab/internal/jwt"
	"github.com/GitCollabCode/GitCollab/microservices/profiles/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func ProfileRouter(p *handlers.Profiles, jwtConf *jwt.GitCollabJwtConf) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentEncoding("application/json"))
	r.Use(handlers.SetContentType("application/json"))
	r.With(p.MiddleWareValidateProfile).Post("/", p.PostProfile)

	r.Get("/search", p.SearchProfile)
	r.Route("/{username}", func(r chi.Router) {
		r.Get("/", p.GetProfile)
		r.Put("/", p.PutProfile)
		r.Patch("/", p.PatchProfile)
		r.Delete("/", p.DeleteProfile)
	})

	r.Route("/skills", func(r chi.Router) {
		r.Use(jwt.JWTBlackList(p.Pd.PDriver))
		r.Use(jwtConf.VerifyJWT(p.Pd.PDriver.Log))
		r.Patch("/", p.PatchSkills)
	})
	return r
}
