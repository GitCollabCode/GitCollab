package gitauth

import (
	"context"
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/jwt"
	"github.com/GitCollabCode/GitCollab/microservices/profiles/data"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type contextKey string

func (c contextKey) String() string { // add custom shit here 😎
	return "mypackage context key " + string(c)
}

var (
	ContextGitClient = contextKey("git-auth-token")
)

func GitClient(p *data.ProfileData) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		/*
		 * Middleware that creates new github clients and passes them through context
		 * You need to use the Verifier middleware first to make sure the user is
		 * actually authenticated! This middleware also relies on some other fields in
		 * context that are supplied by the Verifier middleware.
		 *
		 * To use the client, retrieve and cast from context:
		 *	client := r.Context.Value(gitauth.ContextGitClient).(*github.Client)
		 */
		fn := func(w http.ResponseWriter, r *http.Request) {
			gitID := r.Context().Value(jwt.ContextGitId)
			if gitID == nil {
				p.PDriver.Log.Error("Could not get github ID from context")
			}

			profile, err := p.GetProfile(gitID.(int))
			if err != nil || profile == nil {
				p.PDriver.Log.Errorf("Could not get profile from db: %s", err.Error())
			}

			// create new client
			token := oauth2.Token{AccessToken: profile.GitHubToken}
			ts := oauth2.StaticTokenSource(&token)
			tc := oauth2.NewClient(context.Background(), ts)
			client := github.NewClient(tc)

			// github client passed in context of routes that use this middleware
			r = r.WithContext(context.WithValue(r.Context(), ContextGitClient, client))
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
