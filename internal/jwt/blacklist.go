package jwt

import (
	"context"
	"errors"
	"net/http"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/go-chi/jwtauth"
	"github.com/sirupsen/logrus"
)

type contextKey struct {
	name string
}

var (
	ErrorCtxKey = &contextKey{"Error"}
)

var (
	ErrTokenInvalid = errors.New("token is blacklisted")
	ErrTokenMissing	= errors.New("no token found in header or cookie")
)

// Middleware to check if a given JWT is blacklisted
// All private routes with JWT headers should pass through this
// middleware
func JWTBlackList(ja *jwtauth.JWTAuth, db *db.PostgresDriver, logger *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			jwtVals := [2]string{jwtauth.TokenFromHeader(r), jwtauth.TokenFromCookie(r)}
			var jwtString string
			// iterate over retrieved jwts, check if not null
			for _, jwtString  = range jwtVals {
				if jwtString != "" { // a jwt was
					break
				}
			}

			ctx := r.Context()
			if jwtString == "" { // no jwt found
				ctx = context.WithValue(ctx, ErrorCtxKey, ErrTokenMissing)
				next.ServeHTTP(w, r.WithContext(ctx))
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
			
			for rows.Next()	{
				var retrievedJwt  string
				err := rows.Scan(&retrievedJwt)
				if err != nil {
					logger.Error(err) // TODO: maybe set context
					return
				}
				if retrievedJwt == jwtString{ // found blacklist!
					ctx = context.WithValue(ctx, ErrorCtxKey, ErrTokenInvalid)
				}
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}