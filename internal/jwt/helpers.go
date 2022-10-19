package jwt

import (
	"net/http"
	"strings"
)

type GitCollabJwtConf struct {
	jwtSecret string
}

func NewGitCollabJwtConf(secret string) *GitCollabJwtConf {
	return &GitCollabJwtConf{secret}
}

// Retrieve JWT from the header, return empty string if no JWT
func GetJwtFromHeader(r *http.Request) string {
	// retrieve bearer token from header
	authString := r.Header.Get("Authorization")
	if authString == "" {
		return ""
	}
	splitToken := strings.Split(authString, "Bearer ")
	return splitToken[1]
}
