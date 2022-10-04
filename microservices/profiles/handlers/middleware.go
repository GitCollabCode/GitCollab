package handlers

import (
	"context"
	"net/http"

	"github.com/GitCollabCode/GitCollab/microservices/profiles/data"
)

func SetContentType(contentType string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", contentType)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func (p *Profiles) MiddleWareValidateProfile(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		profile := &data.Profile{}
		err := data.FromJSON(profile, r.Body)
		if err != nil {
			p.log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&ErrorMessage{Message: "Invalid Request: Bad JSON"}, w)
			return
		}

		errs := p.validate.Validate(profile)
		if len(errs) != 0 {
			w.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, w)
			return
		}

		ctx := context.WithValue(r.Context(), ProfileCtx{}, profile)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
