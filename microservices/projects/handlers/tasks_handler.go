package handlers

// import (
// 	"net/http"

// 	jsonio "github.com/GitCollabCode/GitCollab/internal/jsonhttp"
// 	"github.com/GitCollabCode/GitCollab/internal/models"
// 	validate "github.com/GitCollabCode/GitCollab/internal/validator"
// 	"github.com/GitCollabCode/GitCollab/microservices/profiles/data"
// 	profilesModels "github.com/GitCollabCode/GitCollab/microservices/profiles/models"
// 	"github.com/go-chi/chi/v5"
// 	"github.com/jackc/pgx/v5"
// 	"github.com/sirupsen/logrus"
// )

// // Interface for profiles handlers.
// type Tasks struct {
// 	log      *logrus.Logger
// 	validate *validate.Validation
// 	Pd       *data.ProfileData
// }

// // NewProfiles returns initialized Profiles handler struct.
// func NewTasks(log *logrus.Logger, pd *data.ProfileData) *Tasks {
// 	return &Tasks{log, validate.NewValidation(), pd}
// }

// // GetTasks sends all Tasks under a Project.
// func (p *Tasks) GetTasks(w http.ResponseWriter, r *http.Request) {
// 	username := chi.URLParam(r, "username")

// 	profile, err := p.Pd.GetProfileByUsername(username)
// 	if err == pgx.ErrNoRows {
// 		w.WriteHeader(http.StatusNotFound)
// 		err = jsonio.ToJSON(&models.ErrorMessage{Message: "profile does not exist"}, w)
// 		if err != nil {
// 			p.log.Fatalf("GetTasks failed to send error response: %s", err)
// 		}
// 		return
// 	}

// 	if err != nil {
// 		p.log.Errorf("GetTasks database search failed: %s", err.Error())
// 		// NOTE: Repetative code, clean this up
// 		w.WriteHeader(http.StatusInternalServerError)
// 		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
// 		if err != nil {
// 			p.log.Fatalf("GetTasks failed to send error response: %s", err)
// 		}
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")

// 	res := profilesModels.ProfileResp{
// 		Username:  profile.Username,
// 		GithubId:  profile.GitHubUserID,
// 		Email:     profile.Email,
// 		AvatarURL: profile.AvatarURL,
// 		Bio:       profile.Bio,
// 		Skills:    nil, // do this, add to db too
// 		Languages: nil, // do this, add to db too
// 	}

// 	err = jsonio.ToJSON(res, w)
// 	if err != nil {
// 		p.log.Fatalf("GetProfile failed to send response: %s", err)
// 	}
// }

// // PostProfile creates a profile entry in the database, should not be exposed only for testing.
// func (p *Tasks) CreateTask(w http.ResponseWriter, r *http.Request) {
// 	var profileReq profilesModels.ProfileReq

// 	err := p.validate.GetJSON(&profileReq, w, r, p.log)
// 	if err != nil {
// 		p.log.Errorf("SearchProfile failed to decode and validate JSON")
// 		return
// 	}

// 	err = p.Pd.AddProfile(
// 		profileReq.GithubId,
// 		profileReq.GitHubToken,
// 		profileReq.Username,
// 		profileReq.AvatarURL,
// 		profileReq.Email,
// 		profileReq.Bio,
// 	)
// 	if err != nil {
// 		p.log.Errorf("PostProfile database add failed: %s", err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
// 		if err != nil {
// 			p.log.Fatalf("PostProfile failed to send error response: %s", err)
// 		}
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")

// 	err = jsonio.ToJSON(&models.Message{Message: "profile created"}, w)
// 	if err != nil {
// 		p.log.Fatalf("PostProfile failed to send success response: %s", err)
// 	}
// }

// // DeleteProfile deletes a users profile from database.
// func (p *Tasks) DeleteTask(w http.ResponseWriter, r *http.Request) {
// 	// TODO: add a regex check to make sure username follows allowed username format (as middleware maybe?)
// 	username := chi.URLParam(r, "username")

// 	profile, err := p.Pd.GetProfileByUsername(username)
// 	if err == pgx.ErrNoRows {
// 		w.WriteHeader(http.StatusNotFound)
// 		err = jsonio.ToJSON(&models.ErrorMessage{Message: "profile does not exist"}, w)
// 		if err != nil {
// 			p.log.Fatalf("DeleteProfile failed to send error response: %s", err)
// 		}
// 		return
// 	}

// 	if err != nil {
// 		p.log.Errorf("DeleteProfile database search failed: %s", err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
// 		if err != nil {
// 			p.log.Fatalf("DeleteProfile failed to send error response: %s", err)
// 		}
// 		return
// 	}

// 	err = p.Pd.DeleteProfile(profile.GitHubUserID)
// 	if err != nil {
// 		p.log.Errorf("DeleteProfile database delete failed: %s", err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
// 		if err != nil {
// 			p.log.Fatalf("DeleteProfile failed to send error response: %s", err)
// 		}
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/json")

// 	err = jsonio.ToJSON(&models.Message{Message: "profile created"}, w)
// 	if err != nil {
// 		p.log.Fatalf("DeleteProfile failed to send success response: %s", err)
// 	}
// }
