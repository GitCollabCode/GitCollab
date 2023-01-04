package data

import (
	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/GitCollabCode/GitCollab/microservices/profiles/data"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func IsExistingUser(pg *db.PostgresDriver, gitId int, log *logrus.Logger) (*data.Profile, error) {
	pDb := data.NewProfileData(pg)
	profile, err := pDb.GetProfile(gitId)
	if err != nil && err.Error() == pgx.ErrNoRows.Error() {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return profile, nil // user exists, return data
}

func UpdateGitAccessToken(user *data.Profile, token string, pg *db.PostgresDriver, log *logrus.Logger) error {
	// will update the current access token to match that returned by git
	pDb := data.NewProfileData(pg)
	return pDb.UpdateProfileToken(user.GitHubUserID, token)
}

func CreateNewUser(gitId int, gitUser string, gitToken string,
	gitEmail string, gitAvatarUrl string, gitBio string, log *logrus.Logger, pg *db.PostgresDriver) error {
	pDb := data.NewProfileData(pg)
	return pDb.AddProfile(gitId, gitToken, gitUser, gitAvatarUrl, gitEmail, gitBio)
}
