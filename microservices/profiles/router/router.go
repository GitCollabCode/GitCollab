package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func InitRouter(r *chi.Mux, p *handlers.profiles) {
	r.Use(middleware.Logger)

	//use validator middleware to ensure proper json structure is meet

	r.Route("/profile", func(r chi.Router) {
		r.Use(middleware.Recoverer)
		r.Use(middleware.AllowContentEncoding("application/json"))
		r.Use(handlers.SetContentType("application/json"))

		r.Use(cors.Handler(cors.Options{
			//edit later to only accept accept from react app
			AllowedOrigins: []string{"https://*"},
		}))

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
