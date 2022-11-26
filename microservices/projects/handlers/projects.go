package handlers

import (
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/sirupsen/logrus"
)

type ProjectData struct {
	PgConn *db.PostgresDriver
	Log    *logrus.Logger
}

func NewProjectData(pg *db.PostgresDriver, logger *logrus.Logger) *ProjectData {
	return &ProjectData{pg, logger}
}

func (p *ProjectData) GetUserRepos(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *ProjectData) GetRepoInfo(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *ProjectData) GetRepoIssues(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *ProjectData) CreateProject(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *ProjectData) PatchProjectDescription(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *ProjectData) GetUserProjects(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *ProjectData) GetProjectIssues(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
