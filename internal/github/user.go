package githubAPI

import (
	"context"
	"errors"

	"github.com/google/go-github/github"
)

/*
 * Get User ID from users github
 */
func GetUserID(client *github.Client) (int, error) {
	if client == nil {
		return -1, errors.New("could not get client from context")
	}
	user, _, err := client.Users.Get(context.Background(), "")
	return int(*user.ID), err
}

/*
 * Get Github Username from users GitHub
 */
func GetGitUserName(client *github.Client) (string, error) {
	if client == nil {
		return "", errors.New("could not get client from context")
	}
	user, _, err := client.Users.Get(context.Background(), "")
	return *user.Name, err
}

/*
 * Get Icon URL from users github
 */
func GetUserIconURL(client *github.Client) (string, error) {
	if client == nil {
		return "", errors.New("could not get client from context")
	}
	user, _, err := client.Users.Get(context.Background(), "")
	return *user.AvatarURL, err
}

/*
 * Get Bio from users github
 */
func GetUserBio(client *github.Client) (string, error) {
	if client == nil {
		return "", errors.New("could not get client from context")
	}
	user, _, err := client.Users.Get(context.Background(), "")
	return *user.Bio, err
}

// needs to be fixed

/*
 * Get Account name, not user name from users github
// */
//func GetAccountName(r *http.Request) (string, error) {
//	client := GetGitClientFromContext(r)
//	if client == nil {
//		return "", errors.New("could not get client from context")
//	}
//	user, _, err := client.Users.Get(context.Background(), "")
//	return *user.Login, err
//}
//
///*
// * Get number of followers from users github
// */
//func GetNumFollowers(r *http.Request) (int, error) {
//	client := GetGitClientFromContext(r)
//	if client == nil {
//		return -1, errors.New("could not get client from context")
//	}
//	user, _, err := client.Users.Get(context.Background(), "")
//	return *user.Followers, err
//}
//
///*
// * Get number of accounts being followed from users github
// */
//func GetNumFollowing(r *http.Request) (int, error) {
//	client := GetGitClientFromContext(r)
//	if client == nil {
//		return -1, errors.New("could not get client from context")
//	}
//	user, _, err := client.Users.Get(context.Background(), "")
//	return *user.Following, err
//}
//
///*
// * Get Company from users github
// */
//func GetCompany(r *http.Request) (string, error) {
//	client := GetGitClientFromContext(r)
//	if client == nil {
//		return "", errors.New("could not get client from context")
//	}
//	user, _, err := client.Users.Get(context.Background(), "")
//	return *user.Company, err
//}
//
///*
// * Get List of languages that a user has used
// */
//func GetUserLanguages(r *http.Request) ([]string, error) {
//	repoService, err := GetRepoService(r)
//	if err != nil {
//		return nil, err
//	}
//
//	opts := github.RepositoryListOptions{
//		Type: "owner",
//	}
//
//	languages := make(map[string]bool) // list of languages
//	repos, _, err := repoService.List(context.Background(), "", &opts)
//
//	// add languages to list
//	for _, repo := range repos {
//		repoOwner := *repo.Owner.Name
//		repoName := *repo.FullName
//		repoLangs, _, err := repoService.ListLanguages(context.Background(), repoOwner, repoName)
//		if err != nil {
//			return nil, err
//		}
//		// add languages
//		for lang, _ := range repoLangs {
//			if !languages[lang] { // not yet in list of languages
//				languages[lang] = true
//			}
//		}
//	}
//
//	// change languages to list
//	userLanguages := make([]string, 0, len(languages))
//	for k := range languages {
//		userLanguages = append(userLanguages, k)
//	}
//	return userLanguages, err
//}
