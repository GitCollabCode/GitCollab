package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/GitCollabCode/GitCollab/internal/jwt"
	"github.com/GitCollabCode/GitCollab/microservices/authentication/github"
	goGithub "github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
)

// struct to hold info for handlers
type Auth struct {
	PgConn         *db.PostgresDriver
	log            *logrus.Logger
	gitOauthID     string
	gitRedirectUrl string
}

const (
	rUrl = "https://github.com/login/oauth/authorize?scope=user&client_id=%s&redirect_uri=%s"
)

// Expected Http Body for login request
type jsonGitOauth struct {
	Code string // github code
}

func NewAuth(log *logrus.Logger, pg *db.PostgresDriver, oauthID string, redirectUrl string) *Auth {
	// create refrence to new auth struct, hold logger
	return &Auth{pg, log, oauthID, redirectUrl}
}

func (a *Auth) GithubRedirectHandler(w http.ResponseWriter, r *http.Request) {
	// get the redirect url for github, when login button is clicked, this will be returned
	// to the frontend
	redirect := fmt.Sprintf(rUrl, a.gitOauthID, a.gitRedirectUrl)
	w.Write([]byte(redirect))
}

func (a *Auth) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Handler for login, occurs when user returns from github redirect.
	// Will first attempt to get code from request body, if valid retrieve a
	// github access token. If this token is good, go ahead and create a JWT
	// for frontend.

	// TODO: If user does not exist in DB, should create jwt and bring to new
	// 		 user flow.
	a.log.Info("Serving login request")
	dec := json.NewDecoder(r.Body)
	var oauth jsonGitOauth
	err := dec.Decode(&oauth)
	if err != nil {
		a.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get github access token from git with code
	gitAccessToken, err := github.GetGithubAccessToken(oauth.Code)
	if err != nil {
		a.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if !gitAccessToken.Valid() {
		a.log.Error("Invalid Github Access Token!")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// create github client to retrieve user info, store in JWT
	oauthClient := github.GitOauthConfig.Client(context.Background(), gitAccessToken)
	client := goGithub.NewClient(oauthClient)
	username, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		a.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// create a new token for the frontend
	tokenString, err := jwt.CreateGitCollabJwt(*username.Login)
	if err != nil {
		a.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// serve token to frontend
	jsonToken := fmt.Sprintf("{token:%s}", tokenString)
	w.Write([]byte(jsonToken))
}

func (a *Auth) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// add jwt's to the blacklist, these will be picked up by the blacklist
	// middleware and refuse access if found
	jwtString := r.Header["Token"]
	if jwtString == nil {
		// add err for frontend
		a.log.Error("jwt not found in header")
		return
	}
	err := jwt.InsertJwtBlacklist(a.PgConn, jwtString[0])
	if err != nil {
		a.log.Error(err)
		// add error for frontend
	}
}
