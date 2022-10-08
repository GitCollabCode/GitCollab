package handlers

import (
	"fmt"
	"net/http"

	"github.com/GitCollabCode/GitCollab/microservices/profiles/data"
	"github.com/sirupsen/logrus"
)

type ProfileCtx struct{}

type Profiles struct {
	log      *logrus.Logger
	validate *data.Validation
	//db  	 *data.MongoDriver

}

func NewProfiles(log *logrus.Logger) *Profiles {
	return &Profiles{log, data.NewValidation()}
}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProfilePath = fmt.Errorf("invalid path, path should be /profile/[username]")

// GenericError is a generic error message returned by a server
type ErrorMessage struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

func (p *Profiles) GetProfile(w http.ResponseWriter, r *http.Request) {
	// username := chi.URLParam(r, "username")

	// profile, err := p.db.GetProfile(username)
	// if err != nil {
	// 	p.log.Error(err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	models.ToJSON(&models.ErrorMessage{Message: "Failed to fetch user profile"}, w)
	// 	return
	// }

	// if profile.GithubUsername == "" {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	models.ToJSON(&models.ErrorMessage{Message: "User does not exist"}, w)
	// 	return
	// }

	// err = data.ToJSON(profile, w)
	// if err != nil {
	// 	p.log.Error(err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	models.ToJSON(&models.ErrorMessage{Message: "Internal Server Error"}, w)
	// 	return
	// }
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *Profiles) PostProfile(w http.ResponseWriter, r *http.Request) {

	//nProfile := r.Context().Value(Profiles{}).(data.Profile)

	// err := p.db.AddProfile(&nProfile)
	// if err != nil {
	// 	p.log.Error(err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	data.ToJSON(&models.ErrorMessage{Message: "Internal Server Error"}, w)
	// 	return
	// }
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *Profiles) PutProfile(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *Profiles) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *Profiles) PatchProfile(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *Profiles) SearchProfile(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
