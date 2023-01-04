package githubAPI

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

// get an access token from git, using code returned from the frontend
// This token is used for all transactions to GitHub.
func GetGithubAccessToken(code string, oauth oauth2.Config) (*oauth2.Token, error) {
	token, err := oauth.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("I didnt make a token!")
		return nil, err
	}
	return token, nil
}
