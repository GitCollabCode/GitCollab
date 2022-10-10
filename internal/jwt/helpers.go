package jwt

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/GitCollabCode/GitCollab/internal/db"
	goJwt "github.com/golang-jwt/jwt"
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

// Create a new JWT for the frontend (NOT GITHUB). All requests from
// the frontend will contain this JWT, if modified or expired, will
// return error to frontend
func CreateGitCollabJwt(username string, gitID int64) (string, error) {
	// create token, return err and token string
	token := goJwt.New(goJwt.SigningMethodHS256)
	claims := token.Claims.(goJwt.MapClaims)
	claims["exp"] = time.Now().Add(48 * time.Hour).Unix()
	claims["authorized"] = true
	claims["user"] = username
	claims["githubID"] = gitID

	gitCollabSecret := os.Getenv("GITCOLLAB_SECRET") // move somewhere else
	tokenStr, err := token.SignedString([]byte(gitCollabSecret))
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return tokenStr, nil
}

// Get the time a given JWT expires at. Takes the jwt string and returns
// time it expires at as time.Time
func getExpTime(tokenString string) (time.Time, error) {
	token, _, err := new(goJwt.Parser).ParseUnverified(tokenString, goJwt.MapClaims{})
	if err != nil {
		return time.Time{}, err
	}

	claims, ok := token.Claims.(goJwt.MapClaims)
	if !ok {
		return time.Time{}, err
	}

	// Convert time to RFC3339
	expString := fmt.Sprint(claims["exp"])
	expTime, err := time.Parse(time.RFC3339, expString)
	if err != nil {
		return time.Time{}, err
	}
	return expTime, nil
}

// Insert a JWT to the blacklist table. Any requests with a header containing
// this JWT will return an error to the frontend
func InsertJwtBlacklist(pg *db.PostgresDriver, jwtString string) error {
	expTime, err := getExpTime(jwtString)
	if err != nil {
		return err
	}

	// Attempting to insert new jwt
	_, err = pg.Connection.Exec(context.Background(),
		`INSERT INTO jwt_blacklist (jwt, invalidated_time)
		 VALUES ($1, $2) ON CONFLICT (jwt) DO NOTHING`,
		jwtString, expTime)

	if err != nil {
		return fmt.Errorf("failed to add jwt to blacklist")
	}

	return nil
}
