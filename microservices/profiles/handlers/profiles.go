package handlers

import (
	"fmt"
	"net/http"

	jsonio "github.com/GitCollabCode/GitCollab/internal/jsonhttp"
	"github.com/GitCollabCode/GitCollab/internal/jwt"
	"github.com/GitCollabCode/GitCollab/microservices/profiles/data"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type ProfileCtx struct{}

type Profiles struct {
	log      *logrus.Logger
	validate *data.Validation
	Pd       *data.ProfileData
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

type ProfileSearchReq struct {
	Username string `json:"username"`
}

type ProfileGetResponse struct {
	Username  string   `json:"username"`
	GithubId  int      `json:"gitID"`
	Email     string   `json:"email"`
	AvatarURL string   `json:"avatarUrl"`
	Bio       string   `json:"bio"`
	Skills    []string `json:"skills"`
	Languages []string `json:"languages"`
}

type ProfilePatchReq struct {
	Username  string   `json:"username"`
	GithubId  int      `json:"gitID"`
	Email     string   `json:"email"`
	AvatarURL string   `json:"avatarUrl"`
	Skills    []string `json:"skills"`
}

type ProfilesResponse struct {
	Username  string `json:"username"`
	GithubId  int    `json:"gitID"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatarUrl"`
}

func (p *Profiles) GetProfile(w http.ResponseWriter, r *http.Request) {
	// TODO: add a regex check to make sure username follows allowed username format (as middleware maybe?)
	username := chi.URLParam(r, "username")

	profile, err := p.Pd.GetProfileByUsername(username)
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		err = jsonio.ToJSON(&ErrorMessage{Message: "profile does not exist"}, w)
		if err != nil {
			p.log.Errorf("GetProfile failed to convert error response to JSON: %s", err)
		}
		return
	}

	if err != nil {
		// NOTE: Repetative code, clean this up and make sure error messages are descriptive
		p.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&ErrorMessage{Message: "Internal Server Error"}, w)
		if err != nil {
			p.log.Errorf("GetProfile failed to convert error response to JSON: %s", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	res := ProfileGetResponse{
		Username:  profile.Username,
		GithubId:  profile.GitHubUserID,
		Email:     profile.Email,
		AvatarURL: profile.AvatarURL,
		Bio:       profile.Bio,
		Skills:    nil, // do this, add to db too
		Languages: nil, // do this, add to db too
	}

	// send response to frontend
	err = jsonio.ToJSON(res, w)
	if err != nil {
		p.log.Fatalf("GetProfile failed to send response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&ErrorMessage{Message: "Internal Server Error"}, w)
		if err != nil {
			p.log.Errorf("GetProfile failed to convert error response to JSON: %s", err)
		}
		return
	}
}

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
		err = jsonio.ToJSON(&ErrorMessage{Message: "Internal Server Error"}, w)
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
		err = jsonio.ToJSON(&ErrorMessage{Message: "Internal Server Error"}, w)
		if err != nil {
			p.log.Errorf("PostProfile failed to convert error response to JSON: %s", err)
		}
	}
}

func (p *Profiles) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	// TODO: add a regex check to make sure username follows allowed username format (as middleware maybe?)
	username := chi.URLParam(r, "username")

	profile, err := p.Pd.GetProfileByUsername(username)
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		err = jsonio.ToJSON(&ErrorMessage{Message: "profile does not exist"}, w)
		if err != nil {
			p.log.Errorf("DeleteProfile failed to convert error response to JSON: %s", err)
		}
		return
	}

	if err != nil {
		p.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&ErrorMessage{Message: "Internal Server Error"}, w)
		if err != nil {
			p.log.Errorf("DeleteProfile failed to convert error response to JSON: %s", err)
		}
		return
	}

	err = p.Pd.DeleteProfile(profile.GitHubUserID)
	if err != nil {
		p.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&ErrorMessage{Message: "Internal Server Error"}, w)
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
		err = jsonio.ToJSON(&ErrorMessage{Message: "Internal Server Error"}, w)
		if err != nil {
			p.log.Errorf("DeleteProfile failed to convert error response to JSON: %s", err)
		}
		return
	}
}

func (p *Profiles) PutProfile(w http.ResponseWriter, r *http.Request) {
	/*
	 *	NOT NEEDED, LEAVING IN INCASE OF FUTURE USE
	 */
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *Profiles) PatchProfile(w http.ResponseWriter, r *http.Request) {
	var profileReq ProfilePatchReq
	err := jsonio.FromJSON(&profileReq, r.Body)
	if err != nil || profileReq.Username == "" {
		p.log.Errorf("Failed to parse json username: %s", err.Error())
	}
	// TODO: NOTHING TO UPDATE YET
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

func (p *Profiles) SearchProfile(w http.ResponseWriter, r *http.Request) {
	var profileReq ProfileSearchReq
	err := jsonio.FromJSON(&profileReq, r.Body)
	if err != nil {
		p.log.Errorf("Failed to parse json username: %s", err.Error())
	}
	profile, err := p.Pd.GetProfileByUsername(profileReq.Username)
	if err != nil {
		p.log.Errorf("Failed to find user by username: %s", err.Error())
	}

	res := ProfilesResponse{
		Username:  profile.Username,
		GithubId:  profile.GitHubUserID,
		Email:     profile.Email,
		AvatarURL: profile.AvatarURL,
	}
	err = jsonio.ToJSON(&res, w)
	if err != nil {
		p.log.Errorf("GetProfile failed to convert error response to JSON: %s", err)
		w.WriteHeader(http.StatusNotFound) // idlk what error to use, change later on
		// when kevin decides
	}

}

func (p *Profiles) PatchSkills(w http.ResponseWriter, r *http.Request) {
	// used to insert a set of skills, does not replace, only appends
	var profileReq ProfilePatchReq
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
	p.Pd.AddProfileSkills(int(userId), profileReq.Skills...)
}

func (p *Profiles) DeleteSkills(w http.ResponseWriter, r *http.Request) {
	// used to remove a set of skills
	var profileReq ProfilePatchReq
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
	p.Pd.RemoveProfileSkills(int(userId), profileReq.Skills...)
}
