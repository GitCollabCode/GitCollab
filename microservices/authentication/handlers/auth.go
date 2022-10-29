package handlers

import (
	"context"
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/GitCollabCode/GitCollab/internal/github"
	"github.com/GitCollabCode/GitCollab/internal/jwt"
	"github.com/GitCollabCode/GitCollab/microservices/authentication/helpers"
	goGithub "github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

// struct to hold info for handlers
type Auth struct {
	PgConn          *db.PostgresDriver
	Log             *logrus.Logger
	oauth           *oauth2.Config
	gitRedirectUrl  string
	gitCollabSecret string
}

const (
	rUrl = "https://github.com/login/oauth/authorize?scope=user&client_id=%s&redirect_uri=%s"
)

// create refrence to new auth struct
// pg = pinter to db driver
// log = logger
// oConf = config for oauth, holds secret and id
// redirectUrl = redirect for frontend, github brings you back here
func NewAuth(pg *db.PostgresDriver, log *logrus.Logger, oConf *oauth2.Config,
	redirectUrl string, gitCollabSecret string) *Auth {
	return &Auth{pg, log, oConf, redirectUrl, gitCollabSecret}
}

// get the redirect url for github, when login button is clicked, this will be returned
// to the frontend
func (a *Auth) GithubRedirectHandler(w http.ResponseWriter, r *http.Request) {
	res, err := helpers.NewRedirectResponse(rUrl)
	if err != nil {
		a.Log.Panic("Failed to create redirect response: %v", err)
	}

	if helpers.WriteJsonResponse(w, res) != nil {
		a.Log.Panic("Failed to write redirect: %v", err)
	}
}

// Handler for login, occurs when user returns from github redirect.
// Will first attempt to get code from request body, if valid retrieve a
// github access token. If this token is good, go ahead and create a JWT
// for frontend.
// TODO: If user does not exist in DB, should create jwt and bring to new user flow.
func (a *Auth) LoginHandler(w http.ResponseWriter, r *http.Request) {
	a.Log.Info("Serving login request")
	oauthReq, err := helpers.ParseGitOauthRequest(r)
	if err != nil {
		a.Log.Errorf("Request missing code: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get github access token from git with code
	gitAccessToken, err := github.GetGithubAccessToken(oauthReq.Code, *a.oauth)
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
	tokenString, err := helpers.CreateGitCollabJwt(*username.Login, *username.ID, a.gitCollabSecret)
	if err != nil {
		a.Log.Errorf("Faild to create a new jwt: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userInfo, err := helpers.IsExistingUser(a.PgConn, int(*username.ID), a.Log)
	if err != nil {
		a.Log.Fatalf("Failed check if user in db: %s", err.Error())
		return
	}

	var email string
	if username.Email == nil {
		email = ""
	} else {
		email = *username.Email
	}

	if userInfo == nil { // did not find the user, create new account
		err := helpers.CreateNewUser(int(*username.ID), *username.Login, gitAccessToken.AccessToken, email, *username.AvatarURL, a.Log, a.PgConn)
		if err != nil {
			a.Log.Errorf("Failed to create new user: %s", err.Error())
			return
		}
	} else if userInfo.GitHubToken != gitAccessToken.AccessToken {
		// found user, but check if token doesnt match
		err := helpers.UpdateGitAccessToken(userInfo, gitAccessToken.AccessToken, a.PgConn, a.Log)
		if err != nil {
			a.Log.Panic("Failed to update users access token: %s", err.Error())
			return
		}
	}

	res, err := helpers.NewLoginResponse(tokenString, userInfo == nil)
	if err != nil {
		a.Log.Panic("failed to create redirect request: %s", err.Error())
	}

	err = helpers.WriteJsonResponse(w, res)
	if err != nil {
		a.Log.Fatalf("failed to serve jwt to frontend: %s", err.Error())
	}
}

// add jwt's to the blacklist, these will be picked up by the blacklist
// middleware and refuse access if found
func (a *Auth) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	jwtString := jwt.GetJwtFromHeader(r)
	if jwtString == "" {
		// add err for frontendF
		a.Log.Error("jwt not found in header")
		return
	}
	a.Log.Infof("Adding jwt %s to blacklist", jwtString)
	err := helpers.InsertJwtBlacklist(a.PgConn, jwtString)
	if err != nil {
		a.Log.Error("Failed to add jwt to blacklist: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// frontend should look for ok after adding to blacklist
	w.WriteHeader(http.StatusOK)
}
