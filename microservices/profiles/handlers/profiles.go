package handlers

import (
	"fmt"
	"net/http"

	"github.com/GitCollabCode/GitCollab/microservices/profiles/data"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type ProfileCtx struct{}

type Profiles struct {
	log      *logrus.Logger
	validate *data.Validation
	pd       *data.ProfileData
}

func NewProfiles(log *logrus.Logger, pd *data.ProfileData) *Profiles {
	return &Profiles{log, data.NewValidation(), pd}
}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProfilePath = fmt.Errorf("invalid path, path should be /profile/[username]")

// ErrorMessage is a generic error message returned by a server
type ErrorMessage struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

func (p *Profiles) GetProfile(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	profile, err := p.pd.GetProfileByUsername(username)
	if err != nil {
		p.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&ErrorMessage{Message: "Internal Server Error"}, w)
		return
	}

	if profile == nil {
		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&ErrorMessage{Message: "profile does not exist"}, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = data.ToJSON(profile, w) //change the returned profile struct later to NOT include github token
	if err != nil {
		p.log.Fatalf("GetProfile failed to send response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&ErrorMessage{Message: "Internal Server Error"}, w)
		return
	}
}

func (p *Profiles) PostProfile(w http.ResponseWriter, r *http.Request) {
	nProfile := r.Context().Value(ProfileCtx{}).(*data.Profile)

	err := p.pd.AddProfile(nProfile.GitHubUserID,
		nProfile.GitHubToken,
		nProfile.Username,
		nProfile.AvatarURL,
		nProfile.Email)
	if err != nil {
		p.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&ErrorMessage{Message: "Internal Server Error"}, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "profile created"

	data.ToJSON(resp, w)
	if err != nil {
		p.log.Fatalf("PostProfile failed to send success response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&ErrorMessage{Message: "Internal Server Error"}, w)
	}
}

func (p *Profiles) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	profile, err := p.pd.GetProfileByUsername(username)
	if err != nil {
		p.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&ErrorMessage{Message: "Internal Server Error"}, w)
		return
	}

	if profile == nil {
		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&ErrorMessage{Message: "profile does not exist"}, w)
		return
	}

	err = p.pd.DeleteProfile(profile.GitHubUserID)
	if err != nil {
		p.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&ErrorMessage{Message: "Internal Server Error"}, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "profile deleted"

	data.ToJSON(resp, w)
	if err != nil {
		p.log.Fatalf("DeleteProfile failed to send success response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&ErrorMessage{Message: "Internal Server Error"}, w)
	}
}

func (p *Profiles) PutProfile(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *Profiles) PatchProfile(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *Profiles) SearchProfile(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
