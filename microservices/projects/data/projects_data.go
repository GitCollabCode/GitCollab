package data

import (
	"context"

	"github.com/GitCollabCode/GitCollab/internal/db"
)

type ProjectData struct {
	PDriver *db.PostgresDriver
}

func NewProjectData(dbDriver *db.PostgresDriver) *ProjectData {
	return &ProjectData{dbDriver}
}

type Project struct {
	ProjectID            int      `db:"project_id"`
	ProjectOwnerId       string   `db:"project_owner_id"`
	ProjectOwnerUsername string   `db:"project_owner_username"`
	ProjectName          string   `db:"project_name"`
	ProjectURL           string   `db:"project_url"`
	ProjectSkills        []string `json:"project_skills"`
	ProjectDescription   string   `json:"project_description"`
	//DateCreated string `json:"date_created"`
}

func (pd *ProjectData) AddProject(ownerID int, ownerUsername string, projectName string, projectURL string,
	description string, skills []string) error {

	sqlString :=
		"INSERT INTO projects(project_owner_id, project_owner_username, project_name, project_url, project_skills, project_description)" +
			"VALUES($1, $2, $3, $4, $5, $6)"

	_, err := pd.PDriver.Pool.Exec(context.Background(), sqlString, ownerID, ownerUsername, projectName, projectURL, skills, description)
	if err != nil {
		pd.PDriver.Log.Errorf("AddProject database INSERT failed: %s", err.Error())
		return err
	}

	return nil
}

func (pd *ProjectData) UpdateProjectName(projectID int, projectName string) error {
	sqlStatement := "UPDATE projects SET project_name = $1 WHERE project_id = $2"
	return pd.PDriver.TransactOneRow(sqlStatement, projectName, projectID)
}

func (pd *ProjectData) UpdateProjectDescription(projectID int, description string) error {
	sqlStatement := "UPDATE projects SET project_description = $1 WHERE project_id = $2"
	return pd.PDriver.TransactOneRow(sqlStatement, description, projectID)
}

/*
 * TODO, i dont know how to add custom types to list yet....
 */

//func (pd *ProjectData) AddProjectskills(projectID int, skill string) error {
//	sqlStatement := "UPDATE projects SET avatar_url = $1 WHERE github_user_id = $2"
//	return pd.projectsTransactOneRow(sqlStatement, avatarURL, githubUserID)
//}

//func (pd *ProjectData) UpdateProjectskills(projectID int, skill string) error {
//	sqlStatement := "UPDATE projects SET avatar_url = $1 WHERE github_user_id = $2"
//	return pd.projectsTransactOneRow(sqlStatement, avatarURL, githubUserID)
//}

//func (pd *ProjectData) DeleteProjectskills(projectID int, skill string) error {
//	sqlStatement := "DELETE FROM projects SET avatar_url = $1 WHERE github_user_id = $2"
//	return pd.projectsTransactOneRow(sqlStatement, avatarURL, githubUserID)
//}

func (pd *ProjectData) DeleteProject(projectID int) error {
	sqlStatement := "DELETE FROM projects WHERE project_id = $1"
	return pd.PDriver.TransactOneRow(sqlStatement, projectID)
}

func (pd *ProjectData) GetProject(projectID int) (*Project, error) {
	var p Project
	sqlStatement := "SELECT * FROM projects WHERE project_id = $1"
	err := pd.PDriver.QueryRow(sqlStatement, &p, projectID)
	return &p, err
}

func (pd *ProjectData) GetUserProjects(username string) ([]Project, error) {
	var p []Project
	sqlStatement := "SELECT * FROM projects WHERE project_owner_username = $1"
	err := pd.PDriver.QueryRows(sqlStatement, &p, username)
	return p, err
}

func (pd *ProjectData) GetTopNProjects(numProjects int) ([]Project, error) {
	var p []Project
	sqlStatement := "SELECT * from projects LIMIT $1"
	err := pd.PDriver.QueryRows(sqlStatement, &p, numProjects)
	return p, err
}
