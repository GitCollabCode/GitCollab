package authModels

// GitHub OAuth code request
// swagger:parameters githubOAuthRequest
type _ struct {
	// in:body
	// Required: true
	Body GitHubOauthReq
}

type GitHubOauthReq struct {
	// temporary code from GitHub to ensure authenticity of user
	// Example: gho_16C7e42F292c6912E7710c838347Ae178B4a
	// Required: true
	Code string `json:"code"`
}

// GitHub login response
// swagger:response loginResponse
type _ struct {
	// in:body
	Body LoginResp
}

type LoginResp struct {
	// users JWT token
	// Example: {Bearer JWT-TOKEN}
	Token string `json:"Token"`
	// indicates if this is a new users logging in for the first time
	NewUser bool `json:"NewUser"`
}

// Redirect URL response
// swagger:response redirectResponse
type _ struct {
	// in:body
	Body GitHubRedirectResp
}

type GitHubRedirectResp struct {
	// redirect url string.
	// Example: https://github.com/login/oauth/authorize?scope=user&client_id=%s&redirect_uri=%s
	RedirectUrl string `json:"RedirectUrl"`
}
