package handlers

import (
	"fmt"
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/db"
	githubAPI "github.com/GitCollabCode/GitCollab/internal/github"
	jsonio "github.com/GitCollabCode/GitCollab/internal/jsonhttp"
	"github.com/GitCollabCode/GitCollab/internal/jwt"
	"github.com/GitCollabCode/GitCollab/microservices/projects/data"
	projectModels "github.com/GitCollabCode/GitCollab/microservices/projects/models"
	"github.com/sirupsen/logrus"
)

// Interface for Projects handlers
type Projects struct {
	PgConn      *db.PostgresDriver
	ProjectData *data.ProjectData
	Log         *logrus.Logger
	JwtConf     *jwt.GitCollabJwtConf
}

// NewProjects returns initialized Projects handler struct
func NewProjects(db *db.PostgresDriver, p *data.ProjectData, jwtConf *jwt.GitCollabJwtConf, logger *logrus.Logger) *Projects {
	return &Projects{db, p, logger, jwtConf}
}

// retrieve list of github repos associated to a given user
// Request all repos that a user owns on github. Will require valid access token
func (p *Projects) GetUserRepos(w http.ResponseWriter, r *http.Request) {
	client, err := githubAPI.GetGitClientFromContext(r)
	if client == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	if err != nil {
		p.Log.Error(err)
	}

	repos, err := githubAPI.GetUserOwnedRepos(client)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	var repoNames []string
	for _, repo := range repos {
		repoNames = append(repoNames, *repo.Name)
		fmt.Println(*repo.Name)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := projectModels.ReposGetResp{Repos: repoNames}
	err = jsonio.ToJSON(resp, w)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
}

// retrieve information about a given repo
// Retrieve basic info about a repo, including name, descriptiom, contributers
func (p *Projects) GetRepoInfo(w http.ResponseWriter, r *http.Request) {
	client, err := githubAPI.GetGitClientFromContext(r)
	if client == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	if err != nil {
		p.Log.Error(err)
	}
	var repoReq projectModels.RepoInfoReq
	err = jsonio.FromJSON(&repoReq, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	repo, err := githubAPI.GetRepoByName(client, repoReq.RepoName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	contributors, err := githubAPI.GetRepoContributers(client, repo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	var contribList []projectModels.Contributor
	for user, userID := range contributors {
		c := projectModels.Contributor{Username: user, GitID: userID}
		contribList = append(contribList, c)
	}
	//TODO ADD LANGUAGES

	resp := projectModels.RepoInfoResp{Contributors: contribList}
	err = jsonio.ToJSON(&resp, w)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// retrieve list of github repos associated to a given user
// Request all repos that a user owns on github. Will require valid access token
func (p *Projects) GetRepoIssues(w http.ResponseWriter, r *http.Request) {
	client, err := githubAPI.GetGitClientFromContext(r)
	if client == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	if err != nil {
		p.Log.Error(err)
	}
	var repoReq projectModels.RepoInfoReq
	err = jsonio.FromJSON(&repoReq, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	repo, err := githubAPI.GetRepoByName(client, repoReq.RepoName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	issues, err := githubAPI.GetRepoIssues(client, repo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	var issueList []projectModels.RepoIssue
	for _, issue := range issues {
		i := projectModels.RepoIssue{Title: *issue.Title, Url: *issue.URL, State: *issue.State}
		issueList = append(issueList, i)
	}

	resp := projectModels.RepoIssueResp{Issues: issueList}
	err = jsonio.ToJSON(&resp, w)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// retrieve list of github repos associated to a given user
// Request all repos that a user owns on github. Will require valid access token
func (p *Projects) CreateProject(w http.ResponseWriter, r *http.Request) {
	client, err := githubAPI.GetGitClientFromContext(r)
	if client == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	if err != nil {
		p.Log.Error(err)
	}
	var repoReq projectModels.RepoInfoReq
	err = jsonio.FromJSON(&repoReq, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	repo, err := githubAPI.GetRepoByName(client, repoReq.RepoName)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value(jwt.ContextGitId).(int)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	username, ok := r.Context().Value(jwt.ContextKeyUser).(string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = p.ProjectData.AddProject(userId, username, *repo.Name, *repo.URL)
	if err != nil {
		p.Log.Error(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// retrieve list of github repos associated to a given user
// Request all repos that a user owns on github. Will require valid access token
func (p *Projects) GetUserProjects(w http.ResponseWriter, r *http.Request) {
	var req projectModels.UserProjectsReq
	err := jsonio.FromJSON(&req, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var resp projectModels.UserProjectsResp
	projects, err := p.ProjectData.GetUserProjects(req.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// create response containing list of projects
	for _, project := range projects {
		resp.Projects = append(resp.Projects, project.ProjectName)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = jsonio.ToJSON(&resp, w)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// retrieve list of github repos associated to a given user
// Request all repos that a user owns on github. Will require valid access token
func (p *Projects) GetProjectIssues(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
