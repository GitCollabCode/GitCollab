package jwt

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	goJwt "github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

type ctxKey string

const (
	USER_CTX_KEY  ctxKey = "username"
	GITID_CTX_KEY ctxKey = "githubID"
)

func parseToken(tokenString string, secret string) (*goJwt.Token, error) {
	if tokenString == "" {
		return nil, fmt.Errorf("could not find jwt string")
	}
	// parse the token
	token, err := goJwt.Parse(tokenString, func(token *goJwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*goJwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unable to parse JWT")
		}
		return []byte(secret), nil // return secret
	})
	// check if retrieved goJwt.Token was good
	if err != nil {
		return nil, fmt.Errorf("could not parse token")
	}
	return token, nil
}

func VerifyJWT(secret string, logger *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			tokenString := GetJwtFromHeader(r)
			token, err := parseToken(tokenString, secret)
			if err != nil { // could not parse the token!
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok { // claims not retrieved!
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if !token.Valid { // no good! backout
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// set context\

			username := claims["user"].(string)
			gitID := claims["githubID"].(string)

			ctx := context.WithValue(r.Context(), USER_CTX_KEY, username)
			ctx = context.WithValue(ctx, GITID_CTX_KEY, gitID)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
