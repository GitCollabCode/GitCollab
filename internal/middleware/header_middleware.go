package middleware

import "net/http"

func SetContentType(contentType string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", contentType)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
