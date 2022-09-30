package jwt

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

func VerifyJWT(logger *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if r.Header["Token"] != nil { // get token from header
				token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok { // if no
						w.WriteHeader(http.StatusUnauthorized)
						_, err := w.Write([]byte("User unauthorized!"))
						return nil, err // set error?
					}
					return "", nil
				})

				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
				}

				if token.Valid { // we gucci, token is valid frfr
					next.ServeHTTP(w, r)
				}

			} else { // no token in header
				w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write([]byte("You're Unauthorized due to No token in the header"))
				if err != nil {
					return
				}
			}
		}
		return http.HandlerFunc(fn)
	}
}
