package jwt

import (
	"os"
	"time"

	goJwt "github.com/golang-jwt/jwt"
)

// todo, move to other package, not handlers
func CreateGitCollabJwt(username string) (string, error) {
	claims := goJwt.MapClaims{}
	claims["exp"] = time.Now().Add(48 * time.Hour) // todo update this
	claims["authorized"] = true
	claims["user"] = username
	// create token, return err and token string
	token := goJwt.NewWithClaims(goJwt.SigningMethodHS256, claims)
	gitCollabSecret := os.Getenv("GITCOLLAB_SECRET") // check if ""
	return  token.SignedString([]byte(gitCollabSecret))
}