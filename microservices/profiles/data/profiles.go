package data

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

// move to db.go
var ErrProfiletNotFound = fmt.Errorf("row not found")
var ErrProfiletMultipleRowsAffected = fmt.Errorf("more than one, row was affected in a single row operation")
var ErrProfiletMultipleRowsRetunred = fmt.Errorf("more than one, row was returned when was expected")

type ProfileData struct {
	db  *pgx.Conn
	log *logrus.Logger
}

func NewProfileData(db *pgx.Conn, log *logrus.Logger) *ProfileData {
	return &ProfileData{db, log}
}

type Profile struct {
	GitHubUserID string `json:"github_user_id"`
	GitHubToken  string `json:"github_token"`
	Username     string `json:"username"`
	AvatarURL    string `json:"avatar_url"`
	Email        string `json:"email" validate:"email"`
}

type Profiles []*Profile

// move to db.go
func (pd *ProfileData) profilesTransactOneRow(sqlStatement string, args ...string) error {
	tx, err := pd.db.Begin(context.Background())
	if err != nil {
		pd.log.Fatal(err)
	}

	defer tx.Rollback(context.Background())

	res, err := tx.Exec(context.Background(), sqlStatement, args)
	if err != nil {
		pd.log.Errorf("profilesTransactOneRow database EXEC failed: %s", err.Error())
		tx.Rollback(context.Background())
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

	return nil
}

func (pd *ProfileData) profilesGetRow(sqlStatement string, args ...string) (*Profile, error) {
	var p Profile
	rows, err := pd.db.Query(context.Background(), sqlStatement, args)
	if err != nil {
		pd.log.Errorf("profilesGetRow Query failed: %s", err.Error())
		return nil, err
	}

	defer rows.Close()

	count := 0
	for rows.Next() {
		if count > 0 {
			err = ErrProfiletMultipleRowsRetunred
			pd.log.Errorf("profilesGetRow failed: %s", err.Error())
			return nil, err
		}

		// NOTE: pgx does not support putting data in struct directly, if there is a better
		// way please replace this.
		err := rows.Scan(&p.GitHubUserID, &p.GitHubToken, &p.Username, &p.AvatarURL, &p.Email)
		if err != nil {
			pd.log.Errorf("profilesGetRow Scan failed: %s", err.Error())
			return nil, err
		}
		count++
	}
	return &p, nil
}

func (pd *ProfileData) AddProfile(githubUserID string, githubToken string, username string, avatarURL string, email string) error {
	sqlString :=
		"INSERT INTO profiles(github_user_id, github_token, username, avatarURL, email)" +
			"VALUES($1, $2, $3, $4, $5)"

	_, err := pd.db.Exec(context.Background(), sqlString, githubUserID, githubToken, username, avatarURL, email)
	if err != nil {
		pd.log.Errorf("AddProfile database INSERT failed: %s", err.Error())
		return err
	}

	return nil
}

func (pd *ProfileData) UpdateProfileToken(githubUserID string, githubToken string) error {
	sqlStatement := "UPDATE profiles SET github_token = $1 WHERE github_user_id = $2"
	return pd.profilesTransactOneRow(sqlStatement, githubToken, githubUserID)
}

func (pd *ProfileData) UpdateProfileUsername(githubUserID string, username string) error {
	sqlStatement := "UPDATE profiles SET username = $1 WHERE github_user_id = $2"
	return pd.profilesTransactOneRow(sqlStatement, username, githubUserID)
}

func (pd *ProfileData) UpdateProfileAvatarURL(githubUserID string, avatarURL string) error {
	sqlStatement := "UPDATE profiles SET avatar_url = $1 WHERE github_user_id = $2"
	return pd.profilesTransactOneRow(sqlStatement, avatarURL, githubUserID)
}

func (pd *ProfileData) UpdateProfileEmail(githubUserID string, email string) error {
	sqlStatement := "UPDATE profiles SET avatar_url = $1 WHERE github_user_id = $2"
	return pd.profilesTransactOneRow(sqlStatement, email, githubUserID)
}

func (pd *ProfileData) DeleteProfile(githubUserID string) error {
	sqlStatement := "DELETE FROM profiles WHERE github_user_id = $1"
	return pd.profilesTransactOneRow(sqlStatement, githubUserID)
}

func (pd *ProfileData) GetProfile(githubUserID string) (*Profile, error) {
	sqlStatement := "SELECT * FROM profiles WHERE github_user_id = $1"
	return pd.profilesGetRow(sqlStatement, githubUserID)
}

func (pd *ProfileData) GetProfileByUsername(username string) (*Profile, error) {
	sqlStatement := "SELECT * FROM profiles WHERE username = $1"
	return pd.profilesGetRow(sqlStatement, username)
}
