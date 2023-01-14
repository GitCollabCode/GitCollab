package githubAPI

import (
	"errors"
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/gitauth"
	"github.com/google/go-github/github"
)

func GetGitClientFromContext(r *http.Request) (*github.Client, error) {
	client := r.Context().Value(gitauth.ContextGitClient)
	if client == nil { // might be able to remove?
		return nil, errors.New("could not find context from request! Make sure to use GitClient middleware")
	}
	return client.(*github.Client), nil
}
