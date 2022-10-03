package router

import (
	"net/http"

	"github.com/GitCollabCode/GitCollab/microservices/authentication/handlers"
	"github.com/go-chi/chi/v5"
)


func AuthRouter(r chi.Router){
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("auth stuff"))
	})

	// add all routes and handlers below
	r.Post("/signin", handlers.LoginHandler)
}

