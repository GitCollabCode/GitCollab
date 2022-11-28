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

func (p *Projects) GetUserRepos(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *Projects) GetRepoInfo(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *Projects) GetRepoIssues(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *Projects) CreateProject(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *Projects) PatchProjectDescription(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *Projects) GetUserProjects(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *Projects) GetProjectIssues(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
