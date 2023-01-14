package handlers

import (
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/db"
	githubAPI "github.com/GitCollabCode/GitCollab/internal/github"
	jsonio "github.com/GitCollabCode/GitCollab/internal/jsonhttp"
	projectModels "github.com/GitCollabCode/GitCollab/microservices/projects/models"
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
	if client == nil {
		w.WriteHeader(http.StatusNotFound)
	}
	repos, err := githubAPI.GetUserOwnedRepos(client)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	var repoNames []string
	for _, repo := range repos {
		repoNames = append(repoNames, *repo.Name)
	}

	resp := projectModels.ProjectGetResp{Projects: repoNames}
	err = jsonio.ToJSON(resp, w)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
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
