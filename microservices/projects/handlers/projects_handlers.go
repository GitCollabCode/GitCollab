package handlers

import (
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/sirupsen/logrus"
)

type Projects struct {
	PgConn *db.PostgresDriver
	Log    *logrus.Logger
}

func NewProjects(pg *db.PostgresDriver, logger *logrus.Logger) *Projects {
	return &Projects{pg, logger}
}

// swagger:route GET /user-repos/ Get github users repos
//
// retrieve list of github repos associated to a given user
//
// Request all repos that a user owns on github. Will require valid access token
//
// responses:
//
//	200: ThingResponse
func (p *Projects) GetUserRepos(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

// swagger:route GET /repo-info/ Get repo info
//
// retrieve information about a given repo
//
// # Retrieve basic info about a repo, including name, descriptiom, contributers
//
// responses:
//
//	200: ThingResponse
func (p *Projects) GetRepoInfo(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

// swagger:route GET /repo-issues/ Get repos associated to a github user
//
// retrieve list of github repos associated to a given user
//
// Request all repos that a user owns on github. Will require valid access token
//
// responses:
//
//	200: ThingResponse
func (p *Projects) GetRepoIssues(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

// swagger:route GET /user-repos/ Get repos associated to a github user
//
// retrieve list of github repos associated to a given user
//
// Request all repos that a user owns on github. Will require valid access token
//
// responses:
//
//	200: ThingResponse
func (p *Projects) CreateProject(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

// swagger:route GET /user-repos/ Get repos associated to a github user
//
// retrieve list of github repos associated to a given user
//
// Request all repos that a user owns on github. Will require valid access token
//
// responses:
//
//	200: ThingResponse
func (p *Projects) PatchProjectDescription(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

// swagger:route GET /user-repos/ Get repos associated to a github user
//
// retrieve list of github repos associated to a given user
//
// Request all repos that a user owns on github. Will require valid access token
//
// responses:
//
//	200: ThingResponse
func (p *Projects) GetUserProjects(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

// swagger:route GET /user-repos/ Get repos associated to a github user
//
// retrieve list of github repos associated to a given user
//
// Request all repos that a user owns on github. Will require valid access token
//
// responses:
//
//	200: ThingResponse
func (p *Projects) GetProjectIssues(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
