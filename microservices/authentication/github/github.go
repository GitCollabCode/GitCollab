package github

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Represents the response received from Github
type githubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func GetGithubAccessToken(code string) string {
	clientID := os.Getenv("GITHUB_CLIENTID")
	clientSecret := os.Getenv("GITHUB_SECRET")

	requestBodyMap := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          code,
	}
	requestJSON, _ := json.Marshal(requestBodyMap)

	req, reqerr := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(requestJSON),
	)

	if reqerr != nil {
		log.Panic("Request creation failed")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, resperr := http.DefaultClient.Do(req)
	if resperr != nil {
		log.Panic("Request failed")
	}

	respbody, _ := ioutil.ReadAll(resp.Body)

	var ghresp githubAccessTokenResponse
	json.Unmarshal(respbody, &ghresp)
	return ghresp.AccessToken
}
