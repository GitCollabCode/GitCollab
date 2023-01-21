package models

// List GitHub repositories response
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

// User projects response
// swagger:response userProjectsResp
type _ struct {
	// in:body
	// Required: true
	Body UserProjectsResp
}

type UserProjectsResp struct {
	// GitCollab projects
	// Example: ["chicken1", "chicken2"]
	Projects []string `json:"projects"`
}
