package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GitCollabCode/GitCollab/microservices/authentication/github"
	"github.com/go-chi/chi/v5"
)

type gitOauth struct {
	code string
}

func AuthRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("auth stuff"))
	})

	r.Get("/signin", func(w http.ResponseWriter, r *http.Request) { // <home>/auth/signin
		var oauth gitOauth
		err := json.NewDecoder(r.Body).Decode(&oauth)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		gitAccessToken := github.GetGithubAccessToken(oauth.code)
		fmt.Println(gitAccessToken)
	})

	return r

}
