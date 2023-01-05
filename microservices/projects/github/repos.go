package githubAPI

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GitHubUserAPI struct {
	client *github.Client
}

func NewGitHubUserAPI(gitHubToken oauth2.TokenSource) *GitHubUserAPI {
	tc := oauth2.NewClient(context.Background(), gitHubToken)
	return &GitHubUserAPI{github.NewClient(tc)}
}

func (g *GitHubUserAPI) GetReposFromUser() ([]*github.Repository, error) {
	user, _, err := g.client.Users.Get(context.Background(), "")
	id := user.GetID()

	var listOptions github.ListOptions
	listOptions.Page = 50 // Number of pages of results to retrieve.

	repos, _, err := g.client.Apps.ListUserRepos(context.Background(), id, &listOptions)

	if err != nil {
		return nil, err
	}
	return repos, err
}
