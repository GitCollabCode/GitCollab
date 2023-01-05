package jwt

import (
	"context"
	"fmt"
	"net/http"

	profilesData "github.com/GitCollabCode/GitCollab/microservices/profiles/data"
	goJwt "github.com/golang-jwt/jwt"
	"golang.org/x/oauth2"
)

type contextKey string

func (c contextKey) String() string { // add custom shit here 😎
	return "mypackage context key " + string(c)
}

var (
	ContextKeyToken = contextKey("token")
	ContextKeyUser  = contextKey("user")
	ContextGitId    = contextKey("gitid")
)

func (g *GitCollabJwtConf) parseToken(tokenString string) (*goJwt.Token, error) {
	token, err := goJwt.Parse(tokenString, func(token *goJwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*goJwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unable to parse JWT")
		}
		return []byte(g.jwtSecret), nil // return secret
	})
	// check if retrieved goJwt.Token was good
	if err != nil {
		return nil, fmt.Errorf("could not parse token")
	}
	return token, nil
}

func (g *GitCollabJwtConf) VerifyJWT(pd *profilesData.ProfileData) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			tokenString := GetJwtFromHeader(r)

			token, err := g.parseToken(tokenString)
			if err != nil { // could not parse the token!
				pd.PDriver.Log.Errorf("Could not validate token %s", err.Error())
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			claims, ok := token.Claims.(goJwt.MapClaims)
			if !ok { // claims not retrieved!
				pd.PDriver.Log.Error("Could not unpack claims")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if !token.Valid { // no good! backout
				pd.PDriver.Log.Error("Invalid JWT")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			username := claims["user"]
			gitID := claims["githubID"]
			if username == nil || gitID == nil {
				pd.PDriver.Log.Error("Not username or GitID found in jwt")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			tokenStr, err := pd.GetTokenByUserID(gitID.(int))
			if err != nil {
				pd.PDriver.Log.Error("gungignaginagung guan . gungi ginga gungi gunga!")
				return
			}
			gitToken := oauth2.Token{AccessToken: tokenStr}

			// set context, change this bih to use map or sync map, kinda wack how im
			// adding both of thees values
			r = r.WithContext(context.WithValue(r.Context(), ContextKeyUser, username.(string)))
			r = r.WithContext(context.WithValue(r.Context(), ContextGitId, gitID.(float64)))
			r = r.WithContext(context.WithValue(r.Context(), ContextKeyToken, gitToken))
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
