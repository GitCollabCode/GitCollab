package router

import (
	"github.com/GitCollabCode/GitCollab/microservices/profiles/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter(r *chi.Mux, p *handlers.Profiles) {

	r.Route("/profile", func(r chi.Router) {
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
	})
}
