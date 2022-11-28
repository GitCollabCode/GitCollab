package data

import (
	"context"
	"fmt"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/sirupsen/logrus"
)

// move to db.go
var ErrProjectNotFound = fmt.Errorf("row not found")
var ErrProjectMultipleRowsAffected = fmt.Errorf("more than one, row was affected in a single row operation")
var ErrProjectMultipleRowsRetunred = fmt.Errorf("more than one, row was returned when was expected")

type ProjectData struct {
	dbDriver *db.PostgresDriver
	log      *logrus.Logger
}

func NewProjectData(dbDriver *db.PostgresDriver, log *logrus.Logger) *ProjectData {
	return &ProjectData{dbDriver, log}
}

type Project struct {
	ProjectID          int    `json:"project_id"`
	ProjectOwner       string `json:"project_owner"`
	ProjectName        string `json:"project_name"`
	ProjectDescription string `json:"project_description"`
	ProjectSkills      string `json:"project_skills"`
	DateCreated        string `json:"date_created"`
}

// move to db.go
func (pd *ProjectData) projectsTransactOneRow(sqlStatement string, args ...any) error {
	tx, err := pd.dbDriver.Connection.Begin(context.Background())
	if err != nil {
		pd.log.Fatal(err)
	}

	res, err := tx.Exec(context.Background(), sqlStatement, args...)
	if err != nil {
		pd.log.Errorf("projectsTransactOneRow database EXEC failed: %s", err.Error())
		rollbackErr := tx.Rollback(context.Background())
		if rollbackErr != nil {
			pd.log.Fatalf("projectsTransactOneRow rollback failed: %s", rollbackErr.Error())
		}
		return err
	}

	if res.RowsAffected() > 1 {
		err = ErrProjectMultipleRowsAffected
		pd.log.Errorf("projectsTransactOneRow failed: %s", err.Error())
		return err
	}

	err = tx.Commit(context.Background())
	if err != nil {
		pd.log.Fatalf("projectsTransactOneRow commit failed: %s", err.Error())
	}

	return err
}

func (pd *ProjectData) projectGetRow(sqlStatement string, args ...any) (*Project, error) {
	var p Project

	err := pd.dbDriver.Connection.QueryRow(context.Background(), sqlStatement, args...).Scan(&p.ProjectID,
		&p.ProjectOwner, &p.ProjectName, &p.ProjectDescription, &p.ProjectSkills, &p.DateCreated)
	if err != nil {
		pd.log.Errorf("projectGetRow Query failed: %s", err.Error())
		return nil, err
	}

	return &p, nil
}

func (pd *ProjectData) AddProject(ownerID int, projectName string, projectDescription string) error {
	sqlString :=
		"INSERT INTO projects(project_owner, project_name, project_description)" +
			"VALUES($1, $2, $3)"

	_, err := pd.dbDriver.Connection.Exec(context.Background(), sqlString, ownerID, projectName, projectDescription)
	if err != nil {
		pd.log.Errorf("AddProject database INSERT failed: %s", err.Error())
		return err
	}

	return nil
}

func (pd *ProjectData) UpdateProjectName(projectID int, projectName string) error {
	sqlStatement := "UPDATE projects SET project_name = $1 WHERE project_id = $2"
	return pd.projectsTransactOneRow(sqlStatement, projectName, projectID)
}

func (pd *ProjectData) UpdateProjectDescription(projectID int, description string) error {
	sqlStatement := "UPDATE projects SET project_description = $1 WHERE project_id = $2"
	return pd.projectsTransactOneRow(sqlStatement, description, projectID)
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
	return pd.projectsTransactOneRow(sqlStatement, projectID)
}

func (pd *ProjectData) GetProject(projectID int) (*Project, error) {
	sqlStatement := "SELECT * FROM projects WHERE github_user_id = $1"
	return pd.projectGetRow(sqlStatement, projectID)
}
