package projectsModels

import "github.com/google/go-github/github"

type GetReposResp struct {
	Repos []*github.Repository `json:"repos"`
}
