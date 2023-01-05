package authModels

type LoginResponse struct {
	Token   string `json:"Token"`
	NewUser bool   `json:"NewUser"`
}

type GitHubRedirectResponse struct {
	RedirectUrl string `json:"RedirectUrl"`
}

// Expected Http Body for login request to github
type GitOauthRequest struct {
	Code string `json:"code"`
}
