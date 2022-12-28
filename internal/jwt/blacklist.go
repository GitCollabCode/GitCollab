package jwt

import (
	"context"
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/jackc/pgx/v5"
)

// Middleware to check if a given JWT is blacklisted
// All private routes with JWT headers should pass through this
// middleware
func JWTBlackList(db *db.PostgresDriver) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			jwtString := GetJwtFromHeader(r)
			if jwtString == "" {
				db.Log.Info("Empty jwt")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			var jwt string
			db.Log.Infof("The jwt searched is: %s", jwtString)

			err := db.Pool.QueryRow(context.Background(),
				"SELECT jwt FROM jwt_blacklist WHERE jwt='abc';").Scan(&jwt)

			if err != nil && err != pgx.ErrNoRows {
				db.Log.Error(err.Error())
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			db.Log.Infof("retreived jwt is %s\n\n", jwt)

			if jwt == jwtString {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// made it through blacklist, good to continue
			db.Log.Info("Request being served")
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
