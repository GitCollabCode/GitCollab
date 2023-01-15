package projectsModels

// Project get request
// swagger:parameters ReposGetResp
type _ struct {
	// in:body
	// Required: true
	Body ReposGetResp
}

type ReposGetResp struct {
	// Github repositories
	// Example: ["chicken1", "chicken2"]
	// Required: true
	Repos []string `json:"repos"`
}

// Request Repo Info
// swagger:parameters RepoInfoReq
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

type Contributor struct {
	// Github username
	// Example: "robotevan"
	Username string `json:"username"`
	// Github ID
	// Example: 12312312
	GitID int `json:"git_id"`
}

// Get Repo Info
// swagger:parameters RepoInfoResp
type _ struct {
	// in:body
	// Required: true
	Body RepoInfoResp
}

type RepoInfoResp struct {
	// Github repositories
	// Example: ["chicken1", "chicken2"]
	Languages    []string      `json:"languages"`
	Contributors []Contributor `json:"contributors"`
}

type RepoIssue struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	State string `json:"state"`
	// TODO add more info here
}

// Get Repo Issues
// swagger:parameters RepoIssueResp
type _ struct {
	// in:body
	// Required: true
	Body RepoIssueResp
}

type RepoIssueResp struct {
	// Github repositories
	// Example: ["chicken1", "chicken2"]
	Issues []RepoIssue `json:"issues"`
}

// Get Req User Projects
// swagger:parameters UserProjectsReq
type _ struct {
	// in:body
	// Required: true
	Body UserProjectsReq
}

type UserProjectsReq struct {
	// Github repositories
	// Example: ["chicken1", "chicken2"]
	Username string `json:"username"`
}

// Get user projects
// swagger:parameters UserProjectsResp
type _ struct {
	// in:body
	// Required: true
	Body UserProjectsResp
}

type UserProjectsResp struct {
	// Github repositories
	// Example: ["chicken1", "chicken2"]
	Projects []string `json:"projects"`
}
