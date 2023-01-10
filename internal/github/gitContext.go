package githubAPI

import (
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/gitauth"
	"github.com/google/go-github/github"
)

func GetGitClientFromContext(r *http.Request) *github.Client {
	client := r.Context().Value(gitauth.ContextGitClient)
	if client == nil { // might be able to remove?
		return nil
	}
	return client.(*github.Client)
}
