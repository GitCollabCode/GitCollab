package jwt

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

func VerifyJWT(secret string, logger *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			tokenString, ok := r.Header["Token"]
			if !ok { // !ok if [Token] doesnt exist
				//w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write([]byte("You're Unauthorized due to No token in the header"))
				if err != nil {
					return
				}
			} else {
				token, err := jwt.Parse(tokenString[0], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok { // invalid
						w.WriteHeader(http.StatusUnauthorized)
						_, err := w.Write([]byte("User unauthorized!"))
						return nil, err // set error?
					}
					return []byte(secret), nil // return secret
				})

				if err != nil {
					//w.WriteHeader(http.StatusUnauthorized)
					fmt.Print("gay")
				}

				if token.Valid { // we gucci, token is valid frfr
					next.ServeHTTP(w, r)
				}
			}
		}
		return http.HandlerFunc(fn)
	}
}
