package jwt

import (
	"context"
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/go-chi/jwtauth"
	"github.com/sirupsen/logrus"
)

type contextKey struct {
	name string
}


// Middleware to check if a given JWT is blacklisted
// All private routes with JWT headers should pass through this
// middleware
func JWTBlackList(ja *jwtauth.JWTAuth, db *db.PostgresDriver, logger *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			jwtString := r.Header.Get("Token")
			if jwtString == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// TODO: Make sure no sql injection
			rows, err := db.Connection.Query(context.Background(),
				"select jwt from jwt_blacklist where jwt=$1",
				jwtString)
			if err != nil {
				logger.Error(err) // TODO: maybe set context
				return
			}
			defer rows.Close() // todo check above for errors

			for rows.Next() {
				var retrievedJwt string
				err := rows.Scan(&retrievedJwt)
				if err != nil {
					logger.Error(err) // TODO: maybe set context
					return
				}
				for jwtString == retrievedJwt { // found blacklist!
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
