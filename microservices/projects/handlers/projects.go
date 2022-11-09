package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/GitCollabCode/GitCollab/internal/github"
	jsonio "github.com/GitCollabCode/GitCollab/internal/jsonhttp"
	"github.com/GitCollabCode/GitCollab/internal/jwt"
	"github.com/GitCollabCode/GitCollab/microservices/authentication/helpers"
	goGithub "github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"

)

// struct to hold info for handlers
type Project struct {
	PgConn          *db.PostgresDriver
	Log             *logrus.Logger
	oauth           *oauth2.Config

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
func NewProjects(pg *db.PostgresDriver, log *logrus.Logger, oConf *oauth2.Config,
	redirectUrl string, gitCollabSecret string) *Auth {
	return &Auth{pg, log, oConf, redirectUrl, gitCollabSecret}
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
	
}

