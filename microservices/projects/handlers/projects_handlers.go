package handlers

import (
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/db"
	githubAPI "github.com/GitCollabCode/GitCollab/internal/github"
	jsonio "github.com/GitCollabCode/GitCollab/internal/jsonhttp"
	"github.com/GitCollabCode/GitCollab/internal/jwt"
	"github.com/GitCollabCode/GitCollab/internal/models"
	"github.com/GitCollabCode/GitCollab/internal/validator"
	"github.com/GitCollabCode/GitCollab/microservices/projects/data"
	projectModels "github.com/GitCollabCode/GitCollab/microservices/projects/models"
	"github.com/sirupsen/logrus"
)

// Interface for Projects handlers
type Projects struct {
	PgConn      *db.PostgresDriver
	ProjectData *data.ProjectData
	validate    validator.Validation
	Log         *logrus.Logger
}

// NewProjects returns initialized Projects handler struct
func NewProjects(db *db.PostgresDriver, p *data.ProjectData, logger *logrus.Logger) *Projects {
	return &Projects{db, p, *validator.NewValidation(), logger}
}

// GetUserRepos retrieve list of github repos associated to a select user
func (p *Projects) GetUserRepos(w http.ResponseWriter, r *http.Request) {
	// TODO: repetitive code use helper
	client, err := githubAPI.GetGitClientFromContext(r)
	if client == nil {
		p.Log.Warning("GetUserRepos client fetch from context returned nothing!")
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.Log.Fatalf("GetUserRepos failed to send error response: %s", err)
		}
		return
	}

	if err != nil {
		p.Log.Errorf("GetUserRepos client fetch from context failed: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.Log.Fatalf("GetUserRepos failed to send error response: %s", err)
		}
		return
	}

	repos, err := githubAPI.GetUserOwnedRepos(client)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "failed to fetch users repos from github"}, w)
		if err != nil {
			p.Log.Fatalf("GetUserRepos failed to send error response: %s", err)
		}
		return
	}

	var repoNames []string
	for _, repo := range repos {
		repoNames = append(repoNames, *repo.Name)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp := projectModels.ReposGetResp{Repos: repoNames}

	err = jsonio.ToJSON(&resp, w)
	if err != nil {
		p.Log.Fatalf("GetUserRepos failed to send response: %s", err)
	}
}

// GetRepoInfo retrieve information about a given repo such as name, description, contributers
func (p *Projects) GetRepoInfo(w http.ResponseWriter, r *http.Request) {
	client, err := githubAPI.GetGitClientFromContext(r)
	if client == nil {
		p.Log.Warning("GetRepoInfo client fetch from context returned nothing!")
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.Log.Fatalf("GetRepoInfo failed to send error response: %s", err)
		}
		return
	}

	if err != nil {
		p.Log.Errorf("GetRepoInfo client fetch from context failed: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.Log.Fatalf("GetRepoInfo failed to send error response: %s", err)
		}
		return
	}

	var repoReq projectModels.RepoInfoReq

	err = p.validate.GetJSON(&repoReq, w, r, p.Log)
	if err != nil {
		p.Log.Errorf("GetRepoInfo failed to decode and validate JSON")
		return
	}

	repo, err := githubAPI.GetRepoByName(client, repoReq.RepoName)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "failed to fetch repo info from github"}, w)
		if err != nil {
			p.Log.Fatalf("GetRepoInfo failed to send error response: %s", err)
		}
		return
	}

	contributors, err := githubAPI.GetRepoContributers(client, repo)
	if err != nil {
		p.Log.Error("GetRepoInfo unable to fetch repo contributers")
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.Log.Fatalf("GetRepoInfo failed to send error response: %s", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	var contribList []projectModels.Contributor
	for user, userID := range contributors {
		c := projectModels.Contributor{Username: user, GitID: userID}
		contribList = append(contribList, c)
	}

	//TODO: ADD LANGUAGES

	resp := projectModels.RepoInfoResp{Contributors: contribList}

	err = jsonio.ToJSON(&resp, w)
	if err != nil {
		p.Log.Fatalf("GetRepoInfo failed to send response: %s", err)
	}
}

// CreateProject create project based on a select repo
func (p *Projects) CreateProject(w http.ResponseWriter, r *http.Request) {
	client, err := githubAPI.GetGitClientFromContext(r)
	if client == nil {
		p.Log.Warning("CreateProject client fetch from context returned nothing!")
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.Log.Fatalf("CreateProject failed to send error response: %s", err)
		}
		return
	}

	if err != nil {
		p.Log.Errorf("CreateProject client fetch from context failed: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.Log.Fatalf("CreateProject failed to send error response: %s", err)
		}
		return
	}

	var repoReq projectModels.CreateRepoReq

	err = p.validate.GetJSON(&repoReq, w, r, p.Log)
	if err != nil {
		p.Log.Errorf("CreateProject failed to decode and validate JSON")
		return
	}

	userId, ok := r.Context().Value(jwt.ContextGitId).(int)
	if !ok {
		p.Log.Errorf("CreateProject failed to fetch GitHub ID from JWT context")
		w.WriteHeader(http.StatusBadRequest)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "invalid GitHub ID"}, w)
		if err != nil {
			p.Log.Fatalf("CreateProject failed to send error response: %s", err)
		}
		return
	}

	username, ok := r.Context().Value(jwt.ContextKeyUser).(string)
	if !ok {
		p.Log.Errorf("CreateProject failed to fetch username from JWT context")
		w.WriteHeader(http.StatusBadRequest)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "invalid username"}, w)
		if err != nil {
			p.Log.Fatalf("CreateProject failed to send error response: %s", err)
		}
		return
	}

	repo, err := githubAPI.GetRepoByName(client, repoReq.RepoName)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "failed to fetch repo info from github"}, w)
		if err != nil {
			p.Log.Fatalf("CreateProject failed to send error response: %s", err)
		}
		return
	}

	err = p.ProjectData.AddProject(userId, username, *repo.Name, *repo.URL, repoReq.Description, repoReq.Skills)
	if err != nil {
		p.Log.Errorf("CreateProject failed add project: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.Log.Fatalf("CreateProject failed to send error response: %s", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	err = jsonio.ToJSON(&models.Message{Message: "project created"}, w)
	if err != nil {
		p.Log.Fatalf("CreateProject failed to send success response: %s", err)
	}
}

// GetUserProjects retrieve list of github repos associated to a given user
func (p *Projects) GetUserProjects(w http.ResponseWriter, r *http.Request) {
	var req projectModels.UserProjectsReq

	err := p.validate.GetJSON(&req, w, r, p.Log)
	if err != nil {
		p.Log.Errorf("GetUserProjects failed to decode and validate JSON")
		return
	}

	var resp projectModels.UserProjectsResp

	projects, err := p.ProjectData.GetUserProjects(req.Username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "failed to fetch user projects"}, w)
		if err != nil {
			p.Log.Fatalf("CreateProject failed to send error response: %s", err)
		}
		return
	}

	// create response containing list of projects
	for _, project := range projects {
		p := projectModels.ProjectInfo{
			ProjectName:        project.ProjectName,
			ProjectDescription: project.ProjectDescription,
			ProjectOwner:       project.ProjectOwnerUsername,
			ProjectSkills:      project.ProjectSkills,
		}
		resp.Projects = append(resp.Projects, p)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	err = jsonio.ToJSON(&resp, w)
	if err != nil {
		p.Log.Fatalf("GetUserProjects failed to send success response: %s", err)
	}
}

// GetUserProjects retrieve list of projects
func (p *Projects) GetSearchProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := p.ProjectData.GetTopNProjects(10)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "Maged could not find any projects"}, w)
		if err != nil {
			p.Log.Fatalf("GetSearchProjects failed to send error response: %s", err)
		}
		return
	}
	var pResp []projectModels.ProjectInfo
	for _, projectInfo := range projects {
		pResp = append(pResp, projectModels.ProjectInfo{
			ProjectOwner:       projectInfo.ProjectOwnerUsername,
			ProjectName:        projectInfo.ProjectName,
			ProjectDescription: projectInfo.ProjectDescription,
			ProjectSkills:      projectInfo.ProjectSkills,
		})
	}

	err = jsonio.ToJSON(&projectModels.SearchProjectsResp{Projects: pResp}, w)
	if err != nil {
		p.Log.Fatalf("GetUserProjects failed to send success response: %s", err)
	}
}
