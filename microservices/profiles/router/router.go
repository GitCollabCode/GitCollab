package router

//import (
//	"net/http"
//	"time"
//
//	"github.com/go-chi/chi/v5"
//	"github.com/go-chi/chi/v5/middleware"
//	"github.com/go-chi/cors"
//	"github.com/go-chi/httprate"
//)
//
//func InitRouter(p *handlers.Profiles) *chi.Mux {
//	r := chi.NewRouter()
//
//	r.Use(middleware.Logger)
//	r.Use(httprate.LimitByIP(100, 1*time.Minute))
//
//	//r.Use(middleware.Heartbeat("/ping"))
//	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("pong")) //can preform a health check instead
//	})
//
//	//use validator middleware to ensure proper json structure is meet
//
//	r.Route("/profile", func(r chi.Router) {
//		r.Use(middleware.Recoverer)
//		r.Use(middleware.AllowContentEncoding("application/json"))
//		r.Use(handlers.SetContentType("application/json"))
//		r.Use(cors.Handler(cors.Options{
//			//edit later to only accept accept from react app
//			AllowedOrigins: []string{"https://*"},
//		}))
//		//JWT auth can be incorporated here as middleware
//
//		r.With(p.MiddleWareValidateProfile).Post("/", p.PostProfile)
//		r.Get("/search", p.SearchProfile)
//
//		r.Route("/{username}", func(r chi.Router) {
//			r.Get("/", p.GetProfile)
//			r.Put("/", p.PutProfile)
//			r.Patch("/", p.PatchProfile)
//			r.Delete("/", p.DeleteProfile)
//		})
//	})
//
//	r.Mount("/debug", middleware.Profiler())
//
//	return r
//}
//
