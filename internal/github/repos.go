package githubAPI

import (
	"context"
	"errors"

	"github.com/google/go-github/github"
)

/*
 * Get List of repos owned by authenticated user
 */
func GetUserOwnedRepos(client *github.Client) ([]*github.Repository, error) {
	repoService := client.Repositories
	if repoService == nil {
		return nil, errors.New("could not get repo service")
	}
	opts := github.RepositoryListOptions{
		Affiliation: "owner",
		Visibility:  "public",
	}

	repos, _, err := repoService.List(context.Background(), "", &opts)
	return repos, err
}

func GetRepoByName(client *github.Client, repoName string) (*github.Repository, error) {
	repoService := client.Repositories
	owner, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		return nil, err
	}
	repo, _, err := repoService.Get(context.Background(), *owner.Login, repoName)
	return repo, err
}

/*
 * return a map of contributers names and their associated github ID
 */
func GetRepoContributers(client *github.Client, repo *github.Repository) (map[string]int, error) {
	repoService := client.Repositories
	username, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		return nil, errors.New("could not get uasername from client")
	}
	contibutors, _, err := repoService.ListContributors(
		context.Background(), *username.Login,
		*repo.Name, &github.ListContributorsOptions{})

	repoContributors := make(map[string]int)
	for _, contributor := range contibutors {
		repoContributors[*contributor.Login] = int(*contributor.ID)
	}

	return repoContributors, err
}

/*
 * Get a list of issues for a repo
 */
func GetRepoIssues(client *github.Client, repo *github.Repository) ([]*github.Issue, error) {
	if !*repo.HasIssues {
		return nil, nil // no issues
	}
	issueService := client.Issues

	issues, _, err := issueService.ListByRepo(context.Background(), *repo.GetOwner().Login, *repo.Name, &github.IssueListByRepoOptions{})
	return issues, err
}
