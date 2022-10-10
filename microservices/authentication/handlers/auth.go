package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/GitCollabCode/GitCollab/internal/github"
	"github.com/GitCollabCode/GitCollab/internal/jwt"
	goGithub "github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

// struct to hold info for handlers
type Auth struct {
	PgConn         *db.PostgresDriver
	Log            *logrus.Logger
	oauth          *oauth2.Config
	gitRedirectUrl string
}

const (
	rUrl = "https://github.com/login/oauth/authorize?scope=user&client_id=%s&redirect_uri=%s"
)

// Expected Http Body for login request
type jsonGitOauth struct {
	Code string // github code
}

// create refrence to new auth struct
// pg = pinter to db driver
// log = logger
// oConf = config for oauth, holds secret and id
// redirectUrl = redirect for frontend, github brings you back here
func NewAuth(pg *db.PostgresDriver, log *logrus.Logger, oConf *oauth2.Config,
	redirectUrl string) *Auth {

	return &Auth{pg, log, oConf, redirectUrl}
}

// get the redirect url for github, when login button is clicked, this will be returned
// to the frontend
func (a *Auth) GithubRedirectHandler(w http.ResponseWriter, r *http.Request) {

	redirect := fmt.Sprintf(rUrl, a.oauth.ClientID, a.gitRedirectUrl)
	w.Write([]byte(redirect))
}

// Handler for login, occurs when user returns from github redirect.
// Will first attempt to get code from request body, if valid retrieve a
// github access token. If this token is good, go ahead and create a JWT
// for frontend.
// TODO: If user does not exist in DB, should create jwt and bring to new user flow.
func (a *Auth) LoginHandler(w http.ResponseWriter, r *http.Request) {
	a.Log.Info("Serving login request")
	dec := json.NewDecoder(r.Body)
	var authCode jsonGitOauth
	err := dec.Decode(&authCode)
	if err != nil {
		a.Log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get github access token from git with code
	gitAccessToken, err := github.GetGithubAccessToken(authCode.Code, *a.oauth)
	if err != nil {
		a.Log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if !gitAccessToken.Valid() {
		a.Log.Error("Invalid Github Access Token!")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// create github client to retrieve user info, store in JWT
	oauthClient := a.oauth.Client(context.Background(), gitAccessToken)
	client := goGithub.NewClient(oauthClient)
	username, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		a.Log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// create a new token for the frontend
	tokenString, err := jwt.CreateGitCollabJwt(*username.Login, *username.ID)
	if err != nil {
		a.Log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// serve token to frontend
	//jsonToken := fmt.Sprintf("{token:%s}", tokenString) // maybe fix json?
	w.Write([]byte(tokenString))
}

func (a *Auth) TestHandler(w http.ResponseWriter, r *http.Request) {
	a.Log.Info("Trying to run verify\n")
	w.Write([]byte(r.Context().Value("username").(string)))
	//w.Write([]byte("yagaaaa"))
}

// add jwt's to the blacklist, these will be picked up by the blacklist
// middleware and refuse access if found
func (a *Auth) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	jwtString := jwt.GetJwtFromHeader(r)
	if jwtString == "" {
		// add err for frontend
		a.Log.Error("jwt not found in header")
		return
	}
	a.Log.Info("Adding jwt %s to blacklist", jwtString)
	err := jwt.InsertJwtBlacklist(a.PgConn, jwtString)
	if err != nil {
		a.Log.Error(err)
		// add error for frontend
	}
}
