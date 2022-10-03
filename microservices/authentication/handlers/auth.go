package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/GitCollabCode/GitCollab/microservices/authentication/github"
	"github.com/golang-jwt/jwt"
	goGithub "github.com/google/go-github/github"
)

type jsonGitOauth struct {
	Code string
}


// todo, move to other package, not handlers
func createToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(48 * time.Hour) // todo update this
	claims["authorized"] = true
	claims["user"] = username
	// create token, return err and token string
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	gitCollabSecret := os.Getenv("GITCOLLAB_SECRET") // check if ""
	return  token.SignedString([]byte(gitCollabSecret))
}


func LoginHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	var oauth jsonGitOauth
	err := dec.Decode(&oauth) // try to retrieve code from request body 
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get github access token from git with code
	gitAccessToken, err := github.GetGithubAccessToken(oauth.Code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if !gitAccessToken.Valid() {
		w.WriteHeader(http.StatusUnauthorized)
	}

	// create new github client, get info
	// TODO, MOVE THIS TO NEW METHOD, WILL GET INFO AND ADD TO DB! OR FETCH FROM DB
	oauthClient := github.GitOauthConfig.Client(context.Background(), gitAccessToken)
	client := goGithub.NewClient(oauthClient)
	fmt.Println("getting username")
	username, _, err := client.Users.Get(context.Background(), "")
	fmt.Printf("USERNAME %s", *username.Login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenString, err := createToken(*username.Login)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// serve token
	jsonToken := fmt.Sprintf("{token:%s}", tokenString)
	w.Write([]byte(jsonToken))
	
}