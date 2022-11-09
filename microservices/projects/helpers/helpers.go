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
	projectName string,
	description string,
	owner string,
	dateCreated string,
}




//Get all projects that a user is associated with
func GetProjectsForProfile (username string) (Projects, error){
	// Get all project_members 
	//left join projects on pm.project id
	// left join users on pm.user_id
	//Where username = pm.user_id

	_, err = pg.Connection.Exec(context.Background(),
		`SELECT project.project_name, project.date_created , project.description, project.project_id
		FROM project_members
		LEFT JOIN project
			ON projects.project_id = project_members.project_id
		LEFT JOIN profile
			ON profile.user_id = project_members.user_id
		 WHERE ${1}=project_members.user_id `,
		username)
}
