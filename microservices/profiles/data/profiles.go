package data

import (
	"context"

	"github.com/GitCollabCode/GitCollab/internal/db"
)

type ProfileData struct {
	PDriver *db.PostgresDriver
}

func NewProfileData(dbDriver *db.PostgresDriver) *ProfileData {
	return &ProfileData{dbDriver}
}

type Profile struct {
	GitHubUserID int
	GitHubToken  string
	Username     string
	AvatarURL    string
	Email        string
	Bio          string
}

type Profiles []*Profile

func (pd *ProfileData) AddProfile(githubUserID int, githubToken string, username string, avatarURL string, email string, bio string) error {
	sqlString :=
		"INSERT INTO profiles(github_user_id, github_token, username, avatar_url, email, bio)" +
			"VALUES($1, $2, $3, $4, $5)"

	_, err := pd.PDriver.Pool.Exec(context.Background(), sqlString, githubUserID, githubToken, username, avatarURL, email, bio)
	if err != nil {
		pd.PDriver.Log.Errorf("AddProfile database INSERT failed: %s", err.Error())
		return err
	}

	return nil
}

func (pd *ProfileData) UpdateProfileToken(githubUserID int, githubToken string) error {
	sqlStatement := "UPDATE profiles SET github_token = $1 WHERE github_user_id = $2"
	return pd.PDriver.TransactOneRow(sqlStatement, githubToken, githubUserID)
}

func (pd *ProfileData) UpdateProfileUsername(githubUserID int, username string) error {
	sqlStatement := "UPDATE profiles SET username = $1 WHERE github_user_id = $2"
	return pd.PDriver.TransactOneRow(sqlStatement, username, githubUserID)
}

func (pd *ProfileData) UpdateProfileAvatarURL(githubUserID int, avatarURL string) error {
	sqlStatement := "UPDATE profiles SET avatar_url = $1 WHERE github_user_id = $2"
	return pd.PDriver.TransactOneRow(sqlStatement, avatarURL, githubUserID)
}

func (pd *ProfileData) UpdateProfileEmail(githubUserID int, email string) error {
	sqlStatement := "UPDATE profiles SET avatar_url = $1 WHERE github_user_id = $2"
	return pd.PDriver.TransactOneRow(sqlStatement, email, githubUserID)
}

func (pd *ProfileData) UpdateProfileBio(githubUserID int, bio string) error {
	sqlStatement := "UPDATE profiles SET bio = $1 WHERE github_user_id = $2"
	return pd.PDriver.TransactOneRow(sqlStatement, bio, githubUserID)
}

func (pd *ProfileData) DeleteProfile(githubUserID int) error {
	sqlStatement := "DELETE FROM profiles WHERE github_user_id = $1"
	return pd.PDriver.TransactOneRow(sqlStatement, githubUserID)
}

func (pd *ProfileData) GetProfile(githubUserID int) (*Profile, error) {
	var p Profile
	sqlStatement := "SELECT * FROM profiles WHERE github_user_id = $1"
	err := pd.PDriver.QueryRow(sqlStatement, &p, githubUserID)
	return &p, err
}

func (pd *ProfileData) GetProfileByUsername(username string) (*Profile, error) {
	var p Profile
	sqlStatement := "SELECT * FROM profiles WHERE username = $1"
	err := pd.PDriver.QueryRow(sqlStatement, &p, username)
	return &p, err
}
