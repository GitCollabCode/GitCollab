package handlers

import (
	"context"
	"net/http"

	jsonio "github.com/GitCollabCode/GitCollab/internal/jsonhttp"
	"github.com/GitCollabCode/GitCollab/internal/models"
	"github.com/GitCollabCode/GitCollab/internal/validator"
	"github.com/GitCollabCode/GitCollab/microservices/profiles/data"
)

// Validates incoming request json body
func (p *Profiles) MiddleWareValidateProfile(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		profile := &data.Profile{}
		err := jsonio.FromJSON(profile, r.Body)
		if err != nil {
			p.log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			err = jsonio.ToJSON(&models.ErrorMessage{Message: "Invalid Request: Bad JSON"}, w)
			if err != nil {
				p.log.Error(err)
			}
			return
		}

		errs := p.validate.Validate(profile)
		if len(errs) != 0 {
			w.WriteHeader(http.StatusUnprocessableEntity)
			err = jsonio.ToJSON(&validator.ValidationErrorResp{Messages: errs.Errors()}, w)
			if err != nil {
				p.log.Error(err)
			}
			return
		}

		ctx := context.WithValue(r.Context(), ProfileCtx{}, profile)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
