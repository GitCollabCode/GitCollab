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
	ProjectID          int    `json:"project_id"`
	ProjectOwner       string `json:"project_owner"`
	ProjectName        string `json:"project_name"`
	ProjectDescription string `json:"project_description"`
	ProjectSkills      string `json:"project_skills"`
	DateCreated        string `json:"date_created"`
}

func (pd *ProjectData) AddProject(ownerID int, projectName string, projectDescription string) error {
	sqlString :=
		"INSERT INTO projects(project_owner, project_name, project_description)" +
			"VALUES($1, $2, $3)"

	_, err := pd.PDriver.Pool.Exec(context.Background(), sqlString, ownerID, projectName, projectDescription)
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
	sqlStatement := "DELETE FROM projects WHERE projectID = $1"
	return pd.PDriver.TransactOneRow(sqlStatement, projectID)
}

func (pd *ProjectData) GetProject(projectID int) (*Project, error) {
	var p Project
	sqlStatement := "SELECT * FROM projects WHERE github_user_id = $1"
	err := pd.PDriver.QueryRow(sqlStatement, &p, projectID)
	return &p, err
}
