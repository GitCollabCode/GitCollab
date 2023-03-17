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
	// Example: 3584d83530557fdd1f46af8289938c8ef79f9dc5
	// Required: true
	Code string `json:"code" validate:"required"`
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
	NewUser  bool   `json:"NewUser"`
	UserName string `json:"UserName"`
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
