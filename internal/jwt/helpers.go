package jwt

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/golang-jwt/jwt"
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
	return token.SignedString([]byte(gitCollabSecret))
}

func getExpTime(tokenString string) (time.Time, error) {
	// maybe verify? if it isn't valid, no need to add.
	// Or run through verifier middleware
	token, _, err := new(goJwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return time.Time{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return time.Time{}, err
	}

	// get expiry time from token
	var expTime time.Time
	switch expClaim := claims["exp"].(type) {
	case float64:
		expTime = time.Unix(int64(expClaim), 0)
	case json.Number: // prob not needed but need to see
		v, _ := expClaim.Int64()
		expTime = time.Unix(v, 0)
	}
	return expTime, nil
}

func InsertJwtBlacklist(pg *db.PostgresDriver, jwtString string) error {

	_, err := getExpTime(jwtString) // TODO, THIS SHOULD ADD TO DB, look into formating
	if err != nil {
		return err
	}
	// "INSERT INTO %s(jwt) VALUES (%s) ON CONFLICT (jwt) DO NOTHING",
	insertSting := fmt.Sprintf("INSERT INTO jwt_blacklist(jwt) VALUES (%s)", jwtString)

	// Attempting to query DB
	_, err = pg.Connection.Exec(context.Background(), insertSting)
	if err != nil {
		return fmt.Errorf("failed to add jwt to blacklist")
	}

	return nil
}
