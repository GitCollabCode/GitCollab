package jwt

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/golang-jwt/jwt"
	goJwt "github.com/golang-jwt/jwt"
)

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
func CreateGitCollabJwt(username string, gitID string) (string, error) {
	claims := goJwt.MapClaims{}
	claims["exp"] = time.Now().Add(48 * time.Hour) // todo update this
	claims["authorized"] = true
	claims["user"] = username
	claims["githubID"] = gitID
	// create token, return err and token string
	token := goJwt.NewWithClaims(goJwt.SigningMethodHS256, claims)
	gitCollabSecret := os.Getenv("GITCOLLAB_SECRET") // check if ""
	return token.SignedString([]byte(gitCollabSecret))
}

// TODO: Fix, either this or jwt generation putting invalid time
// Get the time a given JWT expires at. Takes the jwt string and returns
// time it expires at as time.Time
func getExpTime(tokenString string) (time.Time, error) {
	token, _, err := new(goJwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return time.Time{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
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
