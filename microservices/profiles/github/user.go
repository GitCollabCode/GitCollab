package githubAPI

import (
	"context"
	"errors"
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/gitauth"
	"github.com/google/go-github/github"
)

func getGitClientFromContext(r *http.Request) *github.Client {
	client := r.Context().Value(gitauth.ContextGitClient)
	if client == nil { // might be able to remove?
		return nil
	}
	// can remove above check if this cast returns nil, and does not
	// cause mem fault
	return client.(*github.Client)
}

/*
 * Get User ID from users github
 */
func GetUserID(r *http.Request) (int, error) {
	client := getGitClientFromContext(r)
	if client == nil {
		return -1, errors.New("could not get client from context")
	}
	user, _, err := client.Users.Get(context.Background(), "")
	return int(*user.ID), err
}

/*
 * Get Github Username from users GitHub
 */
func GetGitUserName(r *http.Request) (string, error) {
	client := getGitClientFromContext(r)
	if client == nil {
		return "", errors.New("could not get client from context")
	}
	user, _, err := client.Users.Get(context.Background(), "")
	return *user.Name, err
}

/*
 * Get Icon URL from users github
 */
func GetUserIconURL(r *http.Request) (string, error) {
	client := getGitClientFromContext(r)
	if client == nil {
		return "", errors.New("could not get client from context")
	}
	user, _, err := client.Users.Get(context.Background(), "")
	return *user.AvatarURL, err
}

/*
 * Get Bio from users github
 */
func GetUserBio(r *http.Request) (string, error) {
	client := getGitClientFromContext(r)
	if client == nil {
		return "", errors.New("could not get client from context")
	}
	user, _, err := client.Users.Get(context.Background(), "")
	return *user.Bio, err
}

/*
 * Get Account name, not user name from users github
 */
func GetAccountName(r *http.Request) (string, error) {
	client := getGitClientFromContext(r)
	if client == nil {
		return "", errors.New("could not get client from context")
	}
	user, _, err := client.Users.Get(context.Background(), "")
	return *user.Login, err
}

/*
 * Get number of followers from users github
 */
func GetNumFollowers(r *http.Request) (int, error) {
	client := getGitClientFromContext(r)
	if client == nil {
		return -1, errors.New("could not get client from context")
	}
	user, _, err := client.Users.Get(context.Background(), "")
	return *user.Followers, err
}

/*
 * Get number of accounts being followed from users github
 */
func GetNumFollowing(r *http.Request) (int, error) {
	client := getGitClientFromContext(r)
	if client == nil {
		return -1, errors.New("could not get client from context")
	}
	user, _, err := client.Users.Get(context.Background(), "")
	return *user.Following, err
}

/*
 * Get Company from users github
 */
func GetCompany(r *http.Request) (string, error) {
	client := getGitClientFromContext(r)
	if client == nil {
		return "", errors.New("could not get client from context")
	}
	user, _, err := client.Users.Get(context.Background(), "")
	return *user.Company, err
}
