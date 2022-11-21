package handlers

import (
	"context"
	"fmt"
	"net/http"

	jsonio "github.com/GitCollabCode/GitCollab/internal/jsonhttp"
	"github.com/GitCollabCode/GitCollab/microservices/profiles/data"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"

)

// struct to hold info for handlers
type Project struct {
	PgConn          *db.PostgresDriver
	Log             *logrus.Logger
}


//Used for getting data for Profiles page
type ProjectsProfileResponse struct {
	Projects   string `json:"Projects"`
	Languages  string  `json:"Languages"`
	Skills	   string  `json:"Skills"`
}

type Projects struct {
	projectId string,
	projectName string,
	description string,
	owner string,
	dateCreated string,
}



// create refrence to new auth struct
// pg = pinter to db driver
// log = logger
// oConf = config for oauth, holds secret and id“
// redirectUrl = redirect for frontend, github brings you back here
func NewProjects(pg *db.PostgresDriver, log *logrus.Logger,
	redirectUrl string, gitCollabSecret string) *Auth {
	return &Project{pg, log}
}



// Handler for Profile's page request for projects
// 1.Will get the Username from jwt then get all projects that the user is on
// 2.Will get all the tasks that the user is on
// 3. Parse out the skills and languages the user has been/currently assigned
/* Return JSON {
		Projects:{},
		Languages:{},
		Skills:{}
	}
		*/
// TODO: 
func (p *Projects) ProjectsForProfileHandler(w http.ResponseWriter, r *http.Request) {
	a.Log.Info("Getting Projects for Profile")

	username := chi.URLParam(r, "username")

	project, err := p.PgConn.GetProjectsForProfile(username)
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		err = jsonio.ToJSON(&ErrorMessage{Message: "profile does not exist"}, w)
		if err != nil {
			p.log.Errorf("GetProfile failed to convert error response to JSON: %s", err)
		}
		return
	}
	
}

