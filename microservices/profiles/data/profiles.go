package data

import (
	"fmt"
)

var ErrProfiletNotFound = fmt.Errorf("Profile not found")

type Profile struct {
	GitHubUserID string `json:"github_user_id"`
	GitHubToken  string `json:"github_token"`
	Username     string `json:"username"`
	//Username     string `json:"username" validate:"github_username"`
	Email string `json:"email" validate:"email"`
	Bio   string `json:"bio"`
}

type Profiles []*Profile

//PQX access functions go here
