package jwt

import (
	"net/http"
	"time"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/jackc/pgx/v5"
)

type blacklistData struct {
	expiryTime time.Time `db:"invalidated_time"`
	jwt        string    `db:"jwt"`
}

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

			var jwtData blacklistData

			err := db.QueryRow("SELECT * FROM jwt_blacklist WHERE jwt=$1", &jwtData, jwtString)

			if err != nil && err.Error() != pgx.ErrNoRows.Error() {
				db.Log.Error(err.Error())
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if jwtData.jwt == jwtString {
				db.Log.Info("Jwt in blacklist")
				db.Log.Info(jwtData.expiryTime)
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
