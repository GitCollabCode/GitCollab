package jwt

import (
	"context"
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/sirupsen/logrus"
)

// Middleware to check if a given JWT is blacklisted
// All private routes with JWT headers should pass through this
// middleware
func JWTBlackList(db *db.PostgresDriver, logger *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			jwtString := GetJwtFromHeader(r)
			if jwtString == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			rows, err := db.Pool.Query(context.Background(),
				"select jwt from jwt_blacklist where jwt=$1",
				jwtString)

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			defer rows.Close() // todo check above for errors

			logger.Errorf("err is %s", err.Error())

			for rows.Next() {
				var retrievedJwt string
				err := rows.Scan(&retrievedJwt)
				logger.Infof("jwt is %s", retrievedJwt)
				logger.Errorf("error is %s", err.Error())

				if err != nil || jwtString == retrievedJwt {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
