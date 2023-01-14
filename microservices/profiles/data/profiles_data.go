package data

import (
	"context"
	"errors"

	"github.com/GitCollabCode/GitCollab/internal/db"
)

type ProfileData struct {
	PDriver *db.PostgresDriver
}

func NewProfileData(dbDriver *db.PostgresDriver) *ProfileData {
	return &ProfileData{dbDriver}
}

type Profile struct {
	GitHubUserID int      `db:"github_user_id"`
	GitHubToken  string   `db:"github_token"`
	Username     string   `db:"username"`
	Email        string   `db:"email"`
	AvatarURL    string   `db:"avatar_url"`
	Bio          string   `db:"bio"`
	Skills       []string `db:"skills"`
	Languages    []string `db:"languages"`
}

type Profiles []*Profile

func (pd *ProfileData) AddProfile(githubUserID int, githubToken string, username string, avatarURL string, email string, bio string) error {
	sqlString :=
		"INSERT INTO profiles(github_user_id, github_token, username, avatar_url, email, bio)" +
			"VALUES($1, $2, $3, $4, $5, $6)"

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

func (pd *ProfileData) AddProfileSkills(githubUserID int, skills ...string) error {
	// TODO: Make sure duplicates dont exist
	validSkills := getValidSkills(skills...)
	if len(validSkills) == 0 {
		return errors.New("no skills in list")
	}
	sqlStatement := "UPDATE profiles SET skills = array_cat(skills, $1) WHERE github_user_id = $2"
	return pd.PDriver.TransactOneRow(sqlStatement, validSkills, githubUserID)
}

func (pd *ProfileData) RemoveProfileSkills(githubUserID int, skills ...string) error {
	for _, skill := range skills {
		sqlStatement := "UPDATE profiles SET skills = array_remove(skills, $1) WHERE github_user_id = $2"
		err := pd.PDriver.TransactOneRow(sqlStatement, skill, githubUserID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (pd *ProfileData) AddProfileLanguages(githubUserID int, languages ...string) error {
	// TODO: Make sure duplicates dont exist
	validLanguages := getValidLanguages(languages...)
	if len(validLanguages) == 0 {
		return errors.New("no languages in list")
	}
	sqlStatement := "UPDATE profiles SET languages = array_cat(languages, $1) WHERE github_user_id = $2"
	return pd.PDriver.TransactOneRow(sqlStatement, validLanguages, githubUserID)
}

func (pd *ProfileData) RemoveProfileLanguages(githubUserID int, languages ...string) error {
	for _, language := range languages {
		sqlStatement := "UPDATE profiles SET languages = array_remove(languages, $1) WHERE github_user_id = $2"
		err := pd.PDriver.TransactOneRow(sqlStatement, language, githubUserID)
		if err != nil {
			return err
		}
	}
	return nil
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
	sqlStatement := "SELECT * FROM profiles WHERE LOWER(username) = $1"
	err := pd.PDriver.QueryRow(sqlStatement, &p, username)
	return &p, err
}
