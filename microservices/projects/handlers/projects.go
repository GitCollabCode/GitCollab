package handlers

import (
	"net/http"

	jsonio "github.com/GitCollabCode/GitCollab/internal/jsonhttp"
	verifier "github.com/GitCollabCode/GitCollab/internal/jwt"
	"github.com/GitCollabCode/GitCollab/microservices/projects/data"
	projectsGitAPI "github.com/GitCollabCode/GitCollab/microservices/projects/github"
	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type Projects struct {
	pd  *data.ProjectData
	log *logrus.Logger
}

// ErrorMessage is a generic error message returned by a server
type ErrorMessage struct {
	Message string `json:"message"`
}

type GetUserReposResponse struct {
	repos []*github.Repository `json:"projects"`
}

type GetRepoID struct {
	repoId int `json:repo_id`
}

type GetProjectOwner struct {
	username int `json:"username"`
}

type GetProjectId struct {
	ProjectId int `json:project_id`
}

type GetProjectInfoResponse struct {
	description string `json:project_description`
	name        string `json:project_name`
}

func AddNewProjects(pd *data.ProjectData, logger *logrus.Logger) *Projects {
	return &Projects{pd, logger}
}

func (p *Projects) CreateProject(w http.ResponseWriter, r *http.Request) {
	var repoID GetRepoID

	err := jsonio.FromJSON(&repoID, r.Body)
	if err != nil {
		p.log.Errorf("Failed to deserialize json body: %s", err.Error())
	}

}

func (p *Projects) GetUserRepos(w http.ResponseWriter, r *http.Request) {

	token := r.Context().Value(verifier.ContextKeyToken).(oauth2.Token)

	ts := oauth2.StaticTokenSource(&token)
	client := projectsGitAPI.NewGitHubUserAPI(ts)

	repos, err := client.GetReposFromUser()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err = jsonio.ToJSON(&ErrorMessage{Message: "Failed to fetch User Repos"}, w)
		if err != nil {
			p.log.Errorf("GetProfile failed to convert error response to JSON: %s", err)
		}
		return
	}

	res := GetUserReposResponse{
		repos: repos,
	}

	// send response to frontend
	err = jsonio.ToJSON(res, w)
	if err != nil {
		p.log.Fatalf("GetUserRepos failed to send response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&ErrorMessage{Message: "Internal Server Error"}, w)
		if err != nil {
			p.log.Errorf("GetUserRepos failed to convert error response to JSON: %s", err)
		}
		return
	}

}

/*
	func (p *Projects) GetUserProjects(w http.ResponseWriter, r *http.Request) {
		var projectOwner GetProjectOwner

		err := jsonio.FromJSON(&projectOwner, r.Body)
		if err != nil {
			p.log.Errorf("Failed to parse json body: %s", err.Error())
		}

		projects, err := p.pd.GetProjectsByProjectOwner(projectOwner.Username)
		if err != nil {
			p.log.Errorf("Failed to find projects by project owner: %s", err.Error())
		}

		res := GetUserReposResponse{
			projects: projects,
		}

		err = jsonio.ToJSON(&res, w)
		if err != nil {
			p.log.Errorf("GetProjects failed to convert error response to JSON: %s", err)
			w.WriteHeader(http.StatusNotFound) // idlk what error to use, change later on
			// when kevin decides
		}

}
*/

/*
	func (p *Projects) GetRepoInfo(w http.ResponseWriter, r *http.Request) {
		var ProjectId GetProjectId

		err := jsonio.FromJSON(&ProjectId, r.Body)
		if err != nil {
			p.log.Errorf("Failed to parse json body: %s", err.Error())
		}

		description, err := p.pd.GetProjDescByOwner(ProjectId.projectId)
		if err != nil {
			p.log.Errorf("Failed to find project description from project owner: %s", err.Error())
		}

		name, err := p.pd.GetProjNameByOwner(ProjectId.projectId)
		if err != nil {
			p.log.Errorf("Failed to find project name from project owner: %s", err.Error())
		}

		res := GetProjectInfoResponse{
			description: description,
			name:        name,
		}

		err = jsonio.ToJSON(&res, w)
		if err != nil {
			p.log.Errorf("Get Repo info failed to convert error response to JSON: %s", err)
			w.WriteHeader(http.StatusNotFound) // idlk what error to use, change later on
			// when kevin decides
		}

}
*/
func (p *Projects) GetRepoIssues(w http.ResponseWriter, r *http.Request) {
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
