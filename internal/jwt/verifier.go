package jwt

import (
	"context"
	"fmt"
	"net/http"

	goJwt "github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

type ctxKey string

const (
	CLAIMS_KEY ctxKey = "username"
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
		fmt.Println(err.Error())
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
				fmt.Println("Couldnt parse i stupodui")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			claims, ok := token.Claims.(goJwt.MapClaims)
			if !ok { // claims not retrieved!
				fmt.Println("claims not worky :(")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if !token.Valid { // no good! backout
				fmt.Println("Not valid!!! boy 7what the hwell")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), CLAIMS_KEY, claims["username"])
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
