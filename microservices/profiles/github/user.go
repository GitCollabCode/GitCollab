package githubAPI

// import (
// 	"context"

// 	"github.com/google/go-github/github"
// 	"golang.org/x/oauth2"
// )

// /*
//  * Get the UserName associated to the account
//  */
// func GetUserName(client *github.Client) (string, error) {
// 	token := ""
// 	tc := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
// 	client := github.NewClient(tc)
// 	//opt := &github.UserListOptions{"Name"}
// 	user, _, err := client.Users.Get(context.Background(), "")
// 	if err == nil {
// 		return nil, err
// 	}
// 	return user, err
// }

// /*
//  * Get Icon URL from users github
//  */
// func GetUserIconURL(client *github.Client) (string, err) {

// 	return nil, nil
// }

// /*
//  * Get Bio from users github
//  */
// func GetUserBio(client *github.Client) (string, error) {
// 	return nil, nil
// }

// /*
//  * Get Account name, not user name from users github
//  */
// func GetAccountName(client *github.Client) (string, err) {
// 	return nil, nil
// }

// /*
//  * Get number of followers from users github
//  */
// func GetNumFollowers(client *github.Client) (int, err) {
// 	return nil, nil
// }

// /*
//  * Get number of accounts being followed from users github
//  */
// func GetNumFollowing(client *github.Client) (int, err) {
// 	return nil, nil
// }
