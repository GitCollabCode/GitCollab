package github

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)


var GitOauthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GITHUB_CLIENTID"), // maybe store?
	ClientSecret: os.Getenv("GITHUB_SECRET"), // maybe store?
	Scopes:       []string{"user:email", "user:name"}, // verify what we need
	Endpoint:     githuboauth.Endpoint,
}

func GetGithubAccessToken(code string) (*oauth2.Token, error) {
	token, err := GitOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		fmt.Println("I didnt make a token!")
		return nil, err
	}
	fmt.Println("I made a token!")
	return token, nil
}
