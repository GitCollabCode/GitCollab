package projectsModels

// swagger:response reposGetResp
type _ struct {
	// in:body
	// Required: true
	Body ReposGetResp
}

type ReposGetResp struct {
	// List of GitHub repository names
	// Example: ["chicken1", "chicken2"]
	Repos []string `json:"repos"`
}

// swagger:RequestCreateRepo
type _ struct {
	// in:body
	// Required: true
	Body CreateRepoReq
}

type CreateRepoReq struct {
	// List of GitHub repository names
	// Example: ["chicken1", "chicken2"]
	RepoName    string   `json:"repo_name"`
	Skills      []string `json:"skills"`
	Description string   `json:"description"`
}

// Repo info request
// swagger:parameters repoInfoReq
type _ struct {
	// in:body
	// Required: true
	Body RepoInfoReq
}

type RepoInfoReq struct {
	// Github repository name
	// Example: "sysc4995"
	// Required: true
	RepoName string `json:"repo_name" validate:"required"`
}

// Repo info response
// swagger:response repoInfoResp
type _ struct {
	// in:body
	// Required: true
	Body RepoInfoResp
}

type RepoInfoResp struct {
	// Github repositories
	// Example: ["chicken1", "chicken2"]
	Languages []string `json:"languages"`
	// List of contributors
	// Example [{"username": "wagwan", "git_id": 1234567}]
	Contributors []Contributor `json:"contributors"`
}

type Contributor struct {
	// Github username
	// Example: "robotevan"
	Username string `json:"username"`
	// Github ID
	// Example: 12312312
	GitID int `json:"git_id"`
}

type RepoIssue struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	State string `json:"state"`
	// TODO add more info here
}

// User projects request
// swagger:parameters userProjectsReq
type _ struct {
	// in:body
	// Required: true
	Body UserProjectsReq
}

type UserProjectsReq struct {
	// Profile username
	// Example: wagwan
	Username string `json:"username" validate:"required"`
}

// search projects response
// swagger:response SearchProjectsResp
type _ struct {
	// in:body
	// Required: true
	Body SearchProjectsResp
}

type ProjectInfo struct {
	ProjectName        string   `json:"project_name"`
	ProjectDescription string   `json:"project_description"`
	ProjectOwner       string   `json:"project_owner"`
	ProjectSkills      []string `json:"project_skills"`
}

type SearchProjectsResp struct {
	// GitCollab projects
	Projects []ProjectInfo `json:"projects"`
}

type ProjectResp struct {
	ProjectID            int      `json:"project_id"`
	ProjectOwnerId       string   `json:"project_owner_id"`
	ProjectOwnerUsername string   `json:"project_owner_username"`
	ProjectName          string   `json:"project_name"`
	ProjectURL           string   `json:"project_url"`
	ProjectSkills        []string `json:"project_skills"`
	ProjectDescription   string   `json:"project_description"`
}

// User projects response
// swagger:response userProjectsResp
type _ struct {
	// in:body
	// Required: true
	Body UserProjectsResp
}

type UserProjectsResp struct {
	// GitCollab projects
	Projects []ProjectInfo `json:"projects"`
}
