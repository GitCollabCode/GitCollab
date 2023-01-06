package handlers

import (
	"fmt"
	"net/http"

	jsonio "github.com/GitCollabCode/GitCollab/internal/jsonhttp"
	"github.com/GitCollabCode/GitCollab/internal/jwt"
	"github.com/GitCollabCode/GitCollab/internal/models"
	validate "github.com/GitCollabCode/GitCollab/internal/validator"
	"github.com/GitCollabCode/GitCollab/microservices/profiles/data"
	profilesModels "github.com/GitCollabCode/GitCollab/microservices/profiles/models"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type ProfileCtx struct{}

// Interface for profiles handlers.
type Profiles struct {
	log      *logrus.Logger
	validate *validate.Validation
	Pd       *data.ProfileData
}

// NewProfiles returns initialized Profiles handler struct.
func NewProfiles(log *logrus.Logger, pd *data.ProfileData) *Profiles {
	return &Profiles{log, validate.NewValidation(), pd}
}

// ErrInvalidProductPath is an error message when the product path is not valid.
var ErrInvalidProfilePath = fmt.Errorf("invalid path, path should be /profile/[username]")

// GetProfile returns profile provided username.
func (p *Profiles) GetProfile(w http.ResponseWriter, r *http.Request) {
	// TODO: add a regex check to make sure username follows allowed username format (as middleware maybe?)
	username := chi.URLParam(r, "username")

	profile, err := p.Pd.GetProfileByUsername(username)
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "profile does not exist"}, w)
		if err != nil {
			p.log.Errorf("GetProfile failed to convert error response to JSON: %s", err)
		}
		return
	}

	if err != nil {
		// NOTE: Repetative code, clean this up and make sure error messages are descriptive
		p.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "Internal Server Error"}, w)
		if err != nil {
			p.log.Errorf("GetProfile failed to convert error response to JSON: %s", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	res := profilesModels.ProfileGetResp{
		Username:  profile.Username,
		GithubId:  profile.GitHubUserID,
		Email:     profile.Email,
		AvatarURL: profile.AvatarURL,
		Bio:       profile.Bio,
		Skills:    nil, // do this, add to db too
		Languages: nil, // do this, add to db too
	}

	err = jsonio.ToJSON(res, w)
	if err != nil {
		p.log.Fatalf("GetProfile failed to send response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "Internal Server Error"}, w)
		if err != nil {
			p.log.Errorf("GetProfile failed to convert error response to JSON: %s", err)
		}
		return
	}
}

// PostProfile creates a profile entry in the database, should not be exposed only for testing.
func (p *Profiles) PostProfile(w http.ResponseWriter, r *http.Request) {
	nProfile := r.Context().Value(ProfileCtx{}).(*data.Profile)

	err := p.Pd.AddProfile(
		nProfile.GitHubUserID,
		nProfile.GitHubToken,
		nProfile.Username,
		nProfile.AvatarURL,
		nProfile.Email,
		nProfile.Bio)
	if err != nil {
		p.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "Internal Server Error"}, w)
		if err != nil {
			p.log.Errorf("PostProfile failed to convert error response to JSON: %s", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "profile created"

	err = jsonio.ToJSON(resp, w)
	if err != nil {
		p.log.Fatalf("PostProfile failed to send success response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "Internal Server Error"}, w)
		if err != nil {
			p.log.Errorf("PostProfile failed to convert error response to JSON: %s", err)
		}
	}
}

// DeleteProfile deletes a users profile from database.
func (p *Profiles) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	// TODO: add a regex check to make sure username follows allowed username format (as middleware maybe?)
	username := chi.URLParam(r, "username")

	profile, err := p.Pd.GetProfileByUsername(username)
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "profile does not exist"}, w)
		if err != nil {
			p.log.Errorf("DeleteProfile failed to convert error response to JSON: %s", err)
		}
		return
	}

	if err != nil {
		p.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "Internal Server Error"}, w)
		if err != nil {
			p.log.Errorf("DeleteProfile failed to convert error response to JSON: %s", err)
		}
		return
	}

	err = p.Pd.DeleteProfile(profile.GitHubUserID)
	if err != nil {
		p.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "Internal Server Error"}, w)
		if err != nil {
			p.log.Errorf("DeleteProfile failed to convert error response to JSON: %s", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "profile deleted"

	err = jsonio.ToJSON(resp, w)
	if err != nil {
		p.log.Fatalf("DeleteProfile failed to send success response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "Internal Server Error"}, w)
		if err != nil {
			p.log.Errorf("DeleteProfile failed to convert error response to JSON: %s", err)
		}
		return
	}
}

// SearchProfile returns a profiles information based on input search parameters.
func (p *Profiles) SearchProfile(w http.ResponseWriter, r *http.Request) {
	var profileReq profilesModels.ProfileSearchReq
	err := jsonio.FromJSON(&profileReq, r.Body)
	if err != nil {
		p.log.Errorf("Failed to parse json username: %s", err.Error())
	}
	profile, err := p.Pd.GetProfileByUsername(profileReq.Username)
	if err != nil {
		p.log.Errorf("Failed to find user by username: %s", err.Error())
	}

	res := profilesModels.ProfilesResp{
		Username:  profile.Username,
		GithubId:  profile.GitHubUserID,
		Email:     profile.Email,
		AvatarURL: profile.AvatarURL,
	}
	err = jsonio.ToJSON(&res, w)
	if err != nil {
		p.log.Errorf("GetProfile failed to convert error response to JSON: %s", err)
		w.WriteHeader(http.StatusNotFound)
	}

}

// PatchSkills insert a set of skills into a profile, does not replace, only appends.
func (p *Profiles) PatchSkills(w http.ResponseWriter, r *http.Request) {
	var profileReq profilesModels.ProfilePatchReq
	err := jsonio.FromJSON(&profileReq, r.Body)
	if err != nil {
		p.log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value(jwt.ContextGitId).(float64)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if p.Pd.AddProfileSkills(int(userId), profileReq.Skills...) != nil {
		p.log.Error(err.Error())
	}
}

// DeleteSkills deletes selected skills from a profile.
func (p *Profiles) DeleteSkills(w http.ResponseWriter, r *http.Request) {
	var profileReq profilesModels.ProfilePatchReq
	err := jsonio.FromJSON(&profileReq, r.Body)
	if err != nil {
		p.log.Error(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, ok := r.Context().Value(jwt.ContextGitId).(float64)

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if p.Pd.RemoveProfileSkills(int(userId), profileReq.Skills...) != nil {
		p.log.Error(err.Error())
	}
}
