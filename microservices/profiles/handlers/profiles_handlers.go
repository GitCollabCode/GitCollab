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
			p.log.Fatalf("GetProfile failed to send error response: %s", err)
		}
		return
	}

	if err != nil {
		p.log.Errorf("GetProfile database search failed: %s", err.Error())
		// NOTE: Repetative code, clean this up
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.log.Fatalf("GetProfile failed to send error response: %s", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	res := profilesModels.Profile{
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
		nProfile.Bio,
	)
	if err != nil {
		p.log.Errorf("PostProfile database add failed: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.log.Fatalf("PostProfile failed to send error response: %s", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	err = jsonio.ToJSON(&models.Message{Message: "profile created"}, w)
	if err != nil {
		p.log.Fatalf("PostProfile failed to send success response: %s", err)
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
			p.log.Fatalf("DeleteProfile failed to send error response: %s", err)
		}
		return
	}

	if err != nil {
		p.log.Errorf("DeleteProfile database search failed: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.log.Fatalf("DeleteProfile failed to send error response: %s", err)
		}
		return
	}

	err = p.Pd.DeleteProfile(profile.GitHubUserID)
	if err != nil {
		p.log.Errorf("DeleteProfile database delete failed: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.log.Fatalf("DeleteProfile failed to send error response: %s", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	err = jsonio.ToJSON(&models.Message{Message: "profile created"}, w)
	if err != nil {
		p.log.Fatalf("DeleteProfile failed to send success response: %s", err)
	}
}

// SearchProfile returns a profiles information based on input search parameters.
func (p *Profiles) SearchProfile(w http.ResponseWriter, r *http.Request) {
	var profileReq profilesModels.ProfileSearchReq

	err := jsonio.FromJSON(&profileReq, r.Body)
	if err != nil {
		p.log.Errorf("SearchProfile Failed to decode json request: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "bad request"}, w)
		if err != nil {
			p.log.Fatalf("SearchProfile failed to send error response: %s", err)
		}
	}

	profile, err := p.Pd.GetProfileByUsername(profileReq.Username)
	if err != nil {
		p.log.Errorf("SearchProfile failed to search user by username: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.log.Fatalf("SearchProfile failed to send error response: %s", err)
		}
	}

	if profile == nil {
		w.WriteHeader(http.StatusNoContent)
		err = jsonio.ToJSON(&models.Message{Message: "no profiles found"}, w)
		if err != nil {
			p.log.Fatalf("SearchProfile failed to send success response: %s", err)
		}
	}

	res := profilesModels.SearchProfilesResp{
		Username:  profile.Username,
		GithubId:  profile.GitHubUserID,
		Email:     profile.Email,
		AvatarURL: profile.AvatarURL,
	}
	err = jsonio.ToJSON(&res, w)
	if err != nil {
		p.log.Errorf("SearchProfile failed to send response: %s", err)
	}
}

// PatchSkills insert a set of skills into a profile, does not replace, only appends.
func (p *Profiles) PatchSkills(w http.ResponseWriter, r *http.Request) {
	var profileReq profilesModels.ProfileSkillsReq

	err := jsonio.FromJSON(&profileReq, r.Body)
	if err != nil {
		p.log.Errorf("PatchSkills Failed to decode JSON request: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "bad request"}, w)
		if err != nil {
			p.log.Fatalf("PatchSkills failed to send error response: %s", err)
		}
	}

	userId, ok := r.Context().Value(jwt.ContextGitId).(int)
	if !ok {
		p.log.Errorf("PatchSkills failed to fetch GitHub ID from JWT context: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.log.Fatalf("PatchSkills failed to send error response: %s", err)
		}
		return
	}

	if p.Pd.AddProfileSkills(int(userId), profileReq.Skills...) != nil {
		p.log.Errorf("PatchSkills failed to append skills to profile: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.log.Fatalf("PatchSkills failed to send error response: %s", err)
		}
		return
	}

	err = jsonio.ToJSON(&models.Message{Message: "skills added"}, w)
	if err != nil {
		p.log.Fatalf("PatchSkills failed to send success response: %s", err)
	}
}

// DeleteSkills deletes selected skills from a profile.
func (p *Profiles) DeleteSkills(w http.ResponseWriter, r *http.Request) {
	var profileReq profilesModels.ProfileSkillsReq

	err := jsonio.FromJSON(&profileReq, r.Body)
	if err != nil {
		p.log.Errorf("DeleteSkills failed to decode JSON request: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "bad request"}, w)
		if err != nil {
			p.log.Fatalf("DeleteSkills failed to send error response: %s", err)
		}
	}

	userId, ok := r.Context().Value(jwt.ContextGitId).(float64)
	if !ok {
		p.log.Errorf("DeleteSkills failed to fetch GitHub ID from JWT context: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.log.Fatalf("DeleteSkills failed to send error response: %s", err)
		}
		return
	}

	if p.Pd.RemoveProfileSkills(int(userId), profileReq.Skills...) != nil {
		p.log.Errorf("DeleteSkills failed to delete skills from profile: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.log.Fatalf("DeleteSkills failed to send error response: %s", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	err = jsonio.ToJSON(&models.Message{Message: "skills removed"}, w)
	if err != nil {
		p.log.Fatalf("DeleteSkills failed to send success response: %s", err)
	}
}

func (p *Profiles) GetSkillList(w http.ResponseWriter, r *http.Request) {
	err := jsonio.ToJSON(&profilesModels.GetSkillListResp{Skills: models.Skill[:]}, w)
	if err != nil {
		p.log.Fatalf("GetSkillsList failed to send skill list response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
