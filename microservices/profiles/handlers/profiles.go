package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type ProfileCtx struct{}

type Profiles struct {
	log *logrus.Logger
	//db  *data.MongoDriver
	v *models.Validator
}

func NewProfiles(log *logrus.Logger, v *models.Validator) *Profiles {
	return &Profiles{log, v}
}

func (p Profiles) GetProfile(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	profile, err := p.db.GetProfile(username)
	if err != nil {
		p.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		models.ToJSON(&models.ErrorMessage{Message: "Failed to fetch user profile"}, w)
		return
	}

	if profile.GithubUsername == "" {
		w.WriteHeader(http.StatusNotFound)
		models.ToJSON(&models.ErrorMessage{Message: "User does not exist"}, w)
		return
	}

	err = models.ToJSON(profile, w)
	if err != nil {
		p.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		models.ToJSON(&models.ErrorMessage{Message: "Internal Server Error"}, w)
		return
	}
}

func (p Profiles) PostProfile(w http.ResponseWriter, r *http.Request) {

	nProfile := r.Context().Value(Profiles{}).(models.Profile)

	err := p.db.AddProfile(&nProfile)
	if err != nil {
		p.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		models.ToJSON(&models.ErrorMessage{Message: "Internal Server Error"}, w)
		return
	}
}

func (p Profiles) PutProfile(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p Profiles) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p Profiles) PatchProfile(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p Profiles) SearchProfile(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
