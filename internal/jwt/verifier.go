package jwt

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func VerifyJWT(logger *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Token")
			if token == "" { // empty token
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
		return http.HandlerFunc(fn)
	}
}
