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

type ProjectData struct {
	dbDriver *db.PostgresDriver
	log      *logrus.Logger
}

func NewProjectData(dbDriver *db.PostgresDriver, log *logrus.Logger) *ProjectData {
	return &ProjectData{dbDriver, log}
}

func (pd *ProjectData) AddProject(ownerID int, projectName string, projectDescription string) error {
	sqlString :=
		"INSERT INTO projects(github_owner_id, project_name, project_description)" +
			"VALUES($1, $2, $3)"

	_, err := pd.dbDriver.Connection.Exec(context.Background(), sqlString, ownerID, projectName, projectDescription)
	if err != nil {
		pd.log.Errorf("AddProject database INSERT failed: %s", err.Error())
		return err
	}

	return nil
}
