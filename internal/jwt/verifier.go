package jwt

import (
	"context"
	"fmt"
	"net/http"

	goJwt "github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

type contextKey string

func (c contextKey) String() string { // add custom shit here 😎
	return "mypackage context key " + string(c)
}

var (
	ContextKeyUser = contextKey("user")
	ContextGitId   = contextKey("gitid")
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

func (g *GitCollabJwtConf) VerifyJWT(logger *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			tokenString := GetJwtFromHeader(r)

			token, err := g.parseToken(tokenString)
			if err != nil { // could not parse the token!
				logger.Errorf("Could not validate token %s", err.Error())
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			claims, ok := token.Claims.(goJwt.MapClaims)
			if !ok { // claims not retrieved!
				logger.Error("Could not unpack claims")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if !token.Valid { // no good! backout
				logger.Error("Invalid JWT")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			username := claims["user"]
			gitId := claims["githubID"]
			if username == nil || gitId == nil {
				logger.Error("Not username or GitID found in jwt")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// set context, change this bih to use map or sync map, kinda wack how im
			// adding both of thees values
			r = r.WithContext(context.WithValue(r.Context(), ContextKeyUser, username.(string)))
			r = r.WithContext(context.WithValue(r.Context(), ContextGitId, gitId.(float64)))
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
