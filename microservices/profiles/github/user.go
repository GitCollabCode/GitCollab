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

/*
 * Get User ID from users github
 */
func (g *GitHubUserAPI) GetUserID() (int, error) {
	user, _, err := g.client.Users.Get(context.Background(), "")
	if err != nil {
		return 0, err
	}
	return int(*user.ID), err
}

/*
 * Get Github Username from users GitHub
 */
func (g *GitHubUserAPI) GetGitUserName() (string, error) {

	user, _, err := g.client.Users.Get(context.Background(), "")
	if err != nil {
		return "", err
	}
	return *user.Name, err
}

/*
 * Get Icon URL from users github
 */
func (g *GitHubUserAPI) GetUserIconURL(client *github.Client) (string, error) {
	user, _, err := g.client.Users.Get(context.Background(), "")
	if err != nil {
		return "", err
	}
	return *user.AvatarURL, err
}

/*
 * Get Bio from users github
 */
func (g *GitHubUserAPI) GetUserBio(client *github.Client) (string, error) {
	user, _, err := g.client.Users.Get(context.Background(), "")
	if err != nil {
		return "", err
	}
	return *user.Bio, err
}

/*
 * Get Account name, not user name from users github
 */
func (g *GitHubUserAPI) GetAccountName(client *github.Client) (string, error) {
	user, _, err := g.client.Users.Get(context.Background(), "")
	if err != nil {
		return "", err
	}
	return *user.Login, err
}

/*
 * Get number of followers from users github
 */
func (g *GitHubUserAPI) GetNumFollowers(client *github.Client) (int, error) {
	user, _, err := g.client.Users.Get(context.Background(), "")
	if err != nil {
		return 0, err
	}
	return *user.Followers, err
}

/*
 * Get number of accounts being followed from users github
 */
func (g *GitHubUserAPI) GetNumFollowing(client *github.Client) (int, error) {
	user, _, err := g.client.Users.Get(context.Background(), "")
	if err != nil {
		return 0, err
	}
	return *user.Following, err
}

/*
 * Get Company from users github
 */
func (g *GitHubUserAPI) GetCompany(client *github.Client) (string, error) {
	user, _, err := g.client.Users.Get(context.Background(), "")
	if err != nil {
		return "", err
	}
	return *user.Company, err
}
