package handlers

import (
	"net/http"

	githubAPI "github.com/GitCollabCode/GitCollab/internal/github"
	jsonio "github.com/GitCollabCode/GitCollab/internal/jsonhttp"
	"github.com/GitCollabCode/GitCollab/internal/models"
	projectsModels "github.com/GitCollabCode/GitCollab/microservices/projects/models"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/sirupsen/logrus"
)

// Interface for Projects handlers
type Projects struct {
	PgConn *db.PostgresDriver
	Log    *logrus.Logger
}

// NewProjects returns initialized Projects handler struct
func NewProjects(pg *db.PostgresDriver, logger *logrus.Logger) *Projects {
	return &Projects{pg, logger}
}

// retrieve list of github repos associated to a given user
// Request all repos that a user owns on github. Will require valid access token
func (p *Projects) GetUserRepos(w http.ResponseWriter, r *http.Request) {
	client := githubAPI.GetGitClientFromContext(r)

	repos, err := githubAPI.GetUserOwnedRepos(client)
	if err != nil {
		p.Log.Errorf("GetUserOwnedRepos API call failed: %s", err.Error())
		// NOTE: Repetative code, clean this up
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.Log.Fatalf("GetProfile failed to send error response: %s", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	res := projectsModels.GetReposResp{
		Repos: repos,
	}

	err = jsonio.ToJSON(res, w)
	if err != nil {
		p.Log.Fatalf("GetProfile failed to send response: %s", err)
	}
}

// retrieve information about a given repo
// Retrieve basic info about a repo, including name, descriptiom, contributers
func (p *Projects) GetRepoInfo(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

// retrieve list of github repos associated to a given user
// Request all repos that a user owns on github. Will require valid access token
func (p *Projects) GetRepoIssues(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

// retrieve list of github repos associated to a given user
// Request all repos that a user owns on github. Will require valid access token
func (p *Projects) CreateProject(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

// retrieve list of github repos associated to a given user
// Request all repos that a user owns on github. Will require valid access token
func (p *Projects) PatchProjectDescription(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

// retrieve list of github repos associated to a given user
// Request all repos that a user owns on github. Will require valid access token
func (p *Projects) GetUserProjects(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

// retrieve list of github repos associated to a given user
// Request all repos that a user owns on github. Will require valid access token
func (p *Projects) GetProjectIssues(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
