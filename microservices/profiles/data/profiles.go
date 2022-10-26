package data

import (
	"context"
	"fmt"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/sirupsen/logrus"
)

// move to db.go
var ErrProfiletNotFound = fmt.Errorf("row not found")
var ErrProfiletMultipleRowsAffected = fmt.Errorf("more than one, row was affected in a single row operation")
var ErrProfiletMultipleRowsRetunred = fmt.Errorf("more than one, row was returned when was expected")

type ProfileData struct {
	dbDriver *db.PostgresDriver
	log      *logrus.Logger
}

func NewProfileData(dbDriver *db.PostgresDriver, log *logrus.Logger) *ProfileData {
	return &ProfileData{dbDriver, log}
}

type Profile struct {
	GitHubUserID int    `json:"github_user_id"`
	GitHubToken  string `json:"github_token"`
	Username     string `json:"username"`
	AvatarURL    string `json:"avatar_url"`
	Email        string `json:"email" validate:"email"`
}

type Profiles []*Profile

// move to db.go
func (pd *ProfileData) profilesTransactOneRow(sqlStatement string, args ...any) error {
	tx, err := pd.dbDriver.Connection.Begin(context.Background())
	if err != nil {
		pd.log.Fatal(err)
	}

	defer func() {
		rollbackErr := tx.Rollback(context.Background())
		if rollbackErr != nil {
			pd.log.Fatalf("profilesTransactOneRow rollback failed: %s", rollbackErr.Error())
		}
	}()

	res, err := tx.Exec(context.Background(), sqlStatement, args...)
	if err != nil {
		pd.log.Errorf("profilesTransactOneRow database EXEC failed: %s", err.Error())
		rollbackErr := tx.Rollback(context.Background())
		if rollbackErr != nil {
			pd.log.Fatalf("profilesTransactOneRow rollback failed: %s", rollbackErr.Error())
		}
		return err
	}

	if res.RowsAffected() > 1 {
		err = ErrProfiletMultipleRowsAffected
		pd.log.Errorf("profilesTransactOneRow failed: %s", err.Error())
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		pd.log.Fatalf("profilesTransactOneRow commit failed: %s", err.Error())
	}

	return err
}

func (pd *ProfileData) profilesGetRow(sqlStatement string, args ...any) (*Profile, error) {
	var p Profile

	err := pd.dbDriver.Connection.QueryRow(context.Background(), sqlStatement, args...).Scan(&p.GitHubUserID, &p.GitHubToken, &p.Username, &p.AvatarURL, &p.Email)
	if err != nil {
		pd.log.Errorf("profilesGetRow Query failed: %s", err.Error())
		return nil, err
	}

	return &p, nil
}

func (pd *ProfileData) AddProfile(githubUserID int, githubToken string, username string, avatarURL string, email string) error {
	sqlString :=
		"INSERT INTO profiles(github_user_id, github_token, username, avatar_url, email)" +
			"VALUES($1, $2, $3, $4, $5)"

	_, err := pd.dbDriver.Connection.Exec(context.Background(), sqlString, githubUserID, githubToken, username, avatarURL, email)
	if err != nil {
		pd.log.Errorf("AddProfile database INSERT failed: %s", err.Error())
		return err
	}

	return nil
}

func (pd *ProfileData) UpdateProfileToken(githubUserID int, githubToken string) error {
	sqlStatement := "UPDATE profiles SET github_token = $1 WHERE github_user_id = $2"
	return pd.profilesTransactOneRow(sqlStatement, githubToken, githubUserID)
}

func (pd *ProfileData) UpdateProfileUsername(githubUserID int, username string) error {
	sqlStatement := "UPDATE profiles SET username = $1 WHERE github_user_id = $2"
	return pd.profilesTransactOneRow(sqlStatement, username, githubUserID)
}

func (pd *ProfileData) UpdateProfileAvatarURL(githubUserID int, avatarURL string) error {
	sqlStatement := "UPDATE profiles SET avatar_url = $1 WHERE github_user_id = $2"
	return pd.profilesTransactOneRow(sqlStatement, avatarURL, githubUserID)
}

func (pd *ProfileData) UpdateProfileEmail(githubUserID int, email string) error {
	sqlStatement := "UPDATE profiles SET avatar_url = $1 WHERE github_user_id = $2"
	return pd.profilesTransactOneRow(sqlStatement, email, githubUserID)
}

func (pd *ProfileData) DeleteProfile(githubUserID int) error {
	sqlStatement := "DELETE FROM profiles WHERE github_user_id = $1"
	return pd.profilesTransactOneRow(sqlStatement, githubUserID)
}

func (pd *ProfileData) GetProfile(githubUserID int) (*Profile, error) {
	sqlStatement := "SELECT * FROM profiles WHERE github_user_id = $1"
	return pd.profilesGetRow(sqlStatement, githubUserID)
}

func (pd *ProfileData) GetProfileByUsername(username string) (*Profile, error) {
	sqlStatement := "SELECT * FROM profiles WHERE username = $1"
	return pd.profilesGetRow(sqlStatement, username)
}
