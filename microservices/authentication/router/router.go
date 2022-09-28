package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GitCollabCode/GitCollab/microservices/authentication/github"
	"github.com/GitCollabCode/GitCollab/microservices/authentication/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type gitOauth struct {
	CODE string
}

var (
	log    *logrus.Logger
	secret string
)

func InitAuth(logger *logrus.Logger, jwtSecret string) {
	log = logger
	secret = jwtSecret
}

func AuthRouter() chi.Router {
	r := chi.NewRouter()

	// base route, idk what do yet
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("auth stuff"))
	})

	// signin for our application, create new jwt and swend back to frontend
	r.Get("/sign-in", func(w http.ResponseWriter, r *http.Request) {
		var oauth gitOauth

		// try to retrieve code from request body
		if err := json.NewDecoder(r.Body).Decode(&oauth); err != nil {
			log.Error("No code present in sign-in request")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// get github access token from git with code
		gitAccessToken := github.GetGithubAccessToken(oauth.CODE)
		if gitAccessToken == "" { // failed to get access token
			log.Error("Failed to retrieve GitHub Access token")
			return // probably should do more here? Update context?
		}

		// create a new jwt for the user, TODO: add more to above checks
		// THIS SHOULD ONLY BE DONE IF USER IS ACUTALLY VALID!!!!
		// create claims

		username := "evan" // get from git
		tokenString, err := handlers.CreateToken(username, secret)
		if err != nil {
			log.Error(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return // probably should do more here? Update context?
		}
		jsonToken := fmt.Sprintf("{token:%s}", tokenString)

		w.Write([]byte(jsonToken))
	})
	return r

}
