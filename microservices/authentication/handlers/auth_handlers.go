package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/db"
	jsonio "github.com/GitCollabCode/GitCollab/internal/jsonhttp"
	"github.com/GitCollabCode/GitCollab/internal/jwt"
	"github.com/GitCollabCode/GitCollab/microservices/authentication/data"
	authModels "github.com/GitCollabCode/GitCollab/microservices/authentication/models"
	goGithub "github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

// Interface for Authentication handlers
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

// NewAuth returns initialized Auth handler struct
func NewAuth(pg *db.PostgresDriver, log *logrus.Logger, oConf *oauth2.Config, redirectUrl string, gitCollabSecret string) *Auth {
	return &Auth{pg, log, oConf, redirectUrl, gitCollabSecret}
}

// NOTE: improve the error return for user for these handlers by using ErrorMessage struct.
// See examples in profiles handlers.

// GithubRedirectHandler get the redirect url for github, when login button is
// clicked, this will be returned to the frontend.
func (a *Auth) GithubRedirectHandler(w http.ResponseWriter, r *http.Request) {
	redirect := fmt.Sprintf(rUrl, a.oauth.ClientID, a.gitRedirectUrl)
	err := jsonio.ToJSON(&authModels.GitHubRedirectResp{RedirectUrl: redirect}, w)
	if err != nil {
		a.Log.Panicf("Failed to create redirect response: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// LoginHandler handler for login, occurs when user returns from github redirect.
// Will first attempt to get code from request body, if valid retrieve a github
// access token. If this token is good, go ahead and create a JWT for frontend.
func (a *Auth) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: If user does not exist in DB, should create jwt and bring to new user flow.
	a.Log.Info("Serving login request")
	var githubCodeRes authModels.GitHubOauthReq
	err := jsonio.FromJSON(&githubCodeRes, r.Body)
	if err != nil {
		a.Log.Errorf("Request missing code: %s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get github access token from git with code
	token, err := a.oauth.Exchange(context.Background(), githubCodeRes.Code)
	if err != nil {
		a.Log.Errorf("Failed to get authentication token: %s", err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if !token.Valid() {
		a.Log.Error("Invalid Github Access Token!")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// create github client to retrieve user info, store in JWT
	oauthClient := a.oauth.Client(context.Background(), token)
	client := goGithub.NewClient(oauthClient)
	username, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		a.Log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// create a new token for the frontend
	tokenString, err := jwt.CreateGitCollabJwt(*username.Login, *username.ID, a.gitCollabSecret)
	if err != nil {
		a.Log.Errorf("Faild to create a new jwt: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userInfo, err := data.IsExistingUser(a.PgConn, int(*username.ID), a.Log)
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

	if userInfo == nil {
		// did not find the user, create new account
		err := data.CreateNewUser(int(*username.ID), *username.Login, token.AccessToken, email, *username.AvatarURL, *username.Bio, a.Log, a.PgConn)
		if err != nil {
			a.Log.Errorf("Failed to create new user: %s", err.Error())
			return
		}
	} else if userInfo.GitHubToken != token.AccessToken {
		// found user, but check if token doesnt match
		err := data.UpdateGitAccessToken(userInfo, token.AccessToken, a.PgConn, a.Log)
		if err != nil {
			a.Log.Panicf("Failed to update users access token: %s", err.Error())
			return
		}
	}

	isNewUser := userInfo == nil
	err = jsonio.ToJSON(&authModels.LoginResp{Token: tokenString, NewUser: isNewUser}, w)
	if err != nil {
		a.Log.Fatalf("failed to serve jwt to frontend: %s", err.Error())
	}
}

// LogoutHandler adds jwt to the blacklist, these will be picked up by the blacklist
// middleware and refuse access if found.
func (a *Auth) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	jwtString := jwt.GetJwtFromHeader(r)
	if jwtString == "" {
		// add err for frontendF
		a.Log.Error("jwt not found in header")
		return
	}
	a.Log.Infof("Adding jwt %s to blacklist", jwtString)
	err := jwt.InsertJwtBlacklist(a.PgConn, jwtString)
	if err != nil {
		a.Log.Errorf("Failed to add jwt to blacklist: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
