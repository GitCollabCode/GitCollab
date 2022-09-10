package jwt

import (
	"net/http"
)

//todo rename keypair to tokenpair
//todo change middleware to wrk on keypair

func MiddleWareVerifySessionAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tokens jwtTokenPair
		tokens.GetLatestTokenPair(r)
		// refresh token expired, require login again
		if !tokens.RefreshToken.Valid {
			w.WriteHeader(http.StatusForbidden)
		}
		// access token invalid, attempt refresh
		if !tokens.AccessToken.Valid {
			// attempt to refresh token
			err := tokens.RefreshAccesstoken(w, r)
			if err != TOKEN_VERIFIED || !tokens.AccessToken.Valid {
				// failed to refresh token
				w.WriteHeader(http.StatusForbidden)
			}
		}
		// successfully rewrote access token
		next.ServeHTTP(w, r)
	})
}
