package helpers

/*
 * All json responses and helpers here
 */

import (
	"encoding/json"
	"net/http"
)

type LoginResponse struct {
	Token   string
	NewUser bool
}

type GitHubRedirectResponse struct {
	RedirectUrl string
}

// Expected Http Body for login request to github
type GitOauthRequest struct {
	Code string // github code
}

func WriteJsonResponse(w http.ResponseWriter, res []byte) error {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(res)
	return err
}

func NewLoginResponse(token string, isNewUser bool) ([]byte, error) {
	res := LoginResponse{token, isNewUser}
	return json.Marshal(res)
}

func NewRedirectResponse(url string) ([]byte, error) {
	res := GitHubRedirectResponse{url}
	return json.Marshal(res)
}

func ParseGitOauthRequest(r *http.Request) (GitOauthRequest, error) {
	var req GitOauthRequest
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&req)
	return req, err
}
