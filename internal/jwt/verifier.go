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
				logger.Error("Majed momento")
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
			username := claims["user"]
			gitId := claims["githubID"]
			if username == nil || gitId == nil {
				fmt.Println("No username or git ID in jwt????? huhhhh")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// set context, change this bih to use map or sync map, kinda wack how im
			// adding both of thees values
			ctx := context.WithValue(r.Context(), ContextKeyUser, username.(string))
			ctx2 := context.WithValue(ctx, ContextGitId, gitId.(float64))
			next.ServeHTTP(w, r.WithContext(ctx2))
		}
		return http.HandlerFunc(fn)
	}
}
