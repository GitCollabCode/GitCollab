package helpers

import (
	"context"
	"fmt"
	"time"

	"github.com/GitCollabCode/GitCollab/internal/db"
	"github.com/GitCollabCode/GitCollab/microservices/profiles/data"
	goJwt "github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type Projects struct {
	projectId string,
	ProjectName string,
	description string,
	dateCreated string,
}

type Profile struct {
	ProjectId int    `json:"project_id"`
	ProjectName  string `json:"project_name"`
	Description     string `json:"description"`
	DateCreated		string `json:"date_created"`
}


var ErrProfiletNotFound = fmt.Errorf("row not found")

func (pd *ProfileData) profilesGetRow(sqlStatement string, args ...any) (*Profile, error) {
	var p Projects

	err := pd.dbDriver.Connection.QueryRow(context.Background(), sqlStatement, args...).Scan(&p.ProjectId, &p.ProjectName, &p.Description, &p.DateCreated)
	if err != nil {
		pd.log.Errorf("projectsGetRow Query failed: %s", err.Error())
		return nil, err
	}

	return &p, nil
}

//Get all projects that a user is associated with
func GetProjectsForProfile (username string) (Projects, error){
	// Get all project_members 
	//left join projects on pm.project id
	// left join users on pm.user_id
	//Where username = pm.user_id

	_, err = pg.Connection.Exec(context.Background(),
		`SELECT project.project_id, project.project_name , project.description, project.date_created
		FROM project_members
		LEFT JOIN project
			ON projects.project_id = project_members.project_id
		LEFT JOIN profile
			ON profile.user_id = project_members.user_id
		 WHERE ${1}=project_members.user_id `,
		username)
}
