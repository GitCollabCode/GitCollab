package projectsModels

// Project get request
// swagger:parameters ProjectGetResp
type _ struct {
	// in:body
	// Required: true
	Body ProjectGetResp
}

type ProjectGetResp struct {
	// Github repositories
	// Example: ["chicken1", "chicken2"]
	// Required: true
	Projects []string `json:"projects"`
}
