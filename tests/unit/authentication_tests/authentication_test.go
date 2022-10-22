package authentication_tests

import (
	"os"
	"testing"

	"github.com/GitCollabCode/GitCollab/microservices/authentication/helpers"
	goJwt "github.com/golang-jwt/jwt"
)

const (
	USERNAME = "TEST"
	GIT_ID   = 123
	SECRET   = "NOT OUR ACTUAL SECRET"
)

var (
	JwtTest string
	err     error
	token   *goJwt.Token
)

func TestMain(m *testing.M) {
	// setup goes here

	//
	code := m.Run()
	// teardown goes here

	//
	os.Exit(code)
}

func TestCreateToken(t *testing.T) {
	// test if token generated can be made
	JwtTest, err = helpers.CreateGitCollabJwt(USERNAME, GIT_ID, SECRET)
	if err != nil || JwtTest == "" {
		t.Errorf("GitCollab JWT not created %q", err)
	}
}

func TestJwtParseUnverified(t *testing.T) {
	// test if token can be parsed, DOES NOT VERIFY
	token, _, err = new(goJwt.Parser).ParseUnverified(JwtTest, goJwt.MapClaims{})
	if token == nil || err != nil {
		t.Errorf("created token could not be parsed! %q", err)
	}
}

func TestJwtUsername(t *testing.T) {
	// test if generated jwt contains correct username
	claims, ok := token.Claims.(goJwt.MapClaims)
	if !ok || claims == nil {
		t.Errorf("Could not get claims from jwt!")
	}
	user := claims["user"].(string)
	if user != USERNAME {
		t.Errorf("jwt's username does not match expected!")
	}
}

func TestJwtGitId(t *testing.T) {
	// test if generated jwt contains correct github id
	claims, ok := token.Claims.(goJwt.MapClaims)
	if !ok || claims == nil {
		t.Errorf("Could not get claims from jwt!")
	}
	gitid := claims["githubID"].(float64)
	if gitid != GIT_ID {
		t.Errorf("jwt's git id does not match expected!")
	}
}
