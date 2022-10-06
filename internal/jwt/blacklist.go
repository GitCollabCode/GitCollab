package jwt

import (
	"context"
	"fmt"
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/sirupsen/logrus"
)

type contextKey struct {
	name string
}

// Middleware to check if a given JWT is blacklisted
// All private routes with JWT headers should pass through this
// middleware
func JWTBlackList(db *db.PostgresDriver, logger *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			jwtString := r.URL.Query().Get("jwt")
			if jwtString == "" {
				w.Write([]byte("not found"))
				return
			}

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
				fmt.Println(retrievedJwt)
				if err != nil {
					logger.Error(err) // TODO: maybe set context
					return
				}
				for jwtString == retrievedJwt { // found blacklist!
					fmt.Println("This shit in the blacklist")
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
			}
			fmt.Println("No blkackist pretty epic")
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
