package data

import (
	"context"
	"time"
)

type Task struct {
	TaskID          int       `db:"task_id"`
	ProjectID       int       `db:"project_id"`
	ProjectName     string    `db:"project_name"`
	TaskStatus      string    `db:"task_status"`
	CompletedByID   int       `db:"completed_by_id"`
	CreatedDate     time.Time `db:"date_created_date"`
	CompletedDate   time.Time `db:"completed_date"`
	TaskTitle       string    `db:"task_title"`
	TaskDescription string    `db:"task_description"`
	Diffictly       int       `db:"diffictly"`
	Priority        int       `db:"priority"`
	Skills          []int     `db:"skills"`
}

func (pd *ProjectData) AddTask(taskID int, projectID int, projectName string, taskStatus string, createdDate time.Time, taskTitle string, taskDescription string, diffictly int, priority int, skills []int) error {
	sqlString :=
		"INSERT INTO tasks(task_id, project_id, project_name, task_status, date_created_date, task_title, task_description, diffictly, priority, skills)" +
			"VALUES($1, $2, $3, $4, $5, $6, $7, $8, 9)"

	_, err := pd.PDriver.Pool.Exec(context.Background(), sqlString, taskID, projectID, projectName, taskStatus, createdDate, taskTitle, taskDescription, diffictly, priority, skills)
	if err != nil {
		pd.PDriver.Log.Errorf("AddTask database INSERT failed: %s", err.Error())
		return err
	}

	return nil
}

func (pd *ProjectData) GetTasks(projectName string) ([]Task, error) {
	var tasks []Task
	sqlStatement := "SELECT * FROM tasks WHERE project_name = $1"
	err := pd.PDriver.QueryRow(sqlStatement, &tasks, projectName)
	return tasks, err
}

func (pd *ProjectData) GetTask(projectName string, taskID int) (*Task, error) {
	var t Task
	sqlStatement := "SELECT * FROM tasks WHERE task_id = $1, project_name = $2"
	err := pd.PDriver.QueryRow(sqlStatement, &t, taskID, projectName)
	return &t, err
}

func (pd *ProjectData) DeleteTask(projectName string, taskID int) error {
	sqlStatement := "DELETE FROM tasks WHERE project_name = $1, task_id = $2"
	return pd.PDriver.TransactOneRow(sqlStatement, projectName, taskID)
}

func (pd *ProjectData) CompleteTask(projectName string, taskID int, completedByID int) error {
	sqlStatement := "UPDATE tasks SET completed_by_id = $1, completed_date = $2 WHERE project_name = $3, task_id = $4"
	return pd.PDriver.TransactOneRow(sqlStatement, completedByID, time.Now(), projectName, taskID)
}

func (pd *ProjectData) CoalesceUpdate(projectName string, taskID int, taskTitle string, taskStatus string,
	description string, diffictly int, priority int, skills []int) error {
	sqlStatement := "UPDATE tasks SET task_title = COALESCE($1, task_title), task_status = COALESCE($2, task_status), " +
		"task_description = COALESCE($3, task_description), diffictly = COALESCE($4, diffictly), priority = COALESCE($5, priority), " +
		"skills = COALESCE($6, skills) WHERE project_name = $7, task_id = $8"
	return pd.PDriver.TransactOneRow(sqlStatement, projectName, taskID, taskTitle, taskStatus, description, diffictly, priority, skills)
}

func (pd *ProjectData) UpdateTaskTitle(projectName string, taskID int, taskTitle string) error {
	sqlStatement := "UPDATE projects SET task_title = $1 WHERE project_name = $2, task_id = $3"
	return pd.PDriver.TransactOneRow(sqlStatement, taskTitle, projectName, taskID)
}

func (pd *ProjectData) UpdateTaskStatus(projectName string, taskID int, taskStatus string) error {
	sqlStatement := "UPDATE projects SET task_title = $1 WHERE project_name = $2, task_id = $3"
	return pd.PDriver.TransactOneRow(sqlStatement, taskStatus, projectName, taskID)
}

func (pd *ProjectData) UpdateTaskDescription(projectName string, taskID int, description string) error {
	sqlStatement := "UPDATE projects SET task_description = $1 WHERE project_name = $2, task_id = $3"
	return pd.PDriver.TransactOneRow(sqlStatement, description, projectName, taskID)
}

func (pd *ProjectData) UpdateTaskDiffictly(projectName string, taskID int, diffictly string) error {
	sqlStatement := "UPDATE projects SET diffictly = $1 WHERE project_name = $2, task_id = $3"
	return pd.PDriver.TransactOneRow(sqlStatement, diffictly, projectName, taskID)
}

func (pd *ProjectData) UpdateTaskPriority(projectName string, taskID int, priority string) error {
	sqlStatement := "UPDATE projects SET priority = $1 WHERE project_name = $2, task_id = $3"
	return pd.PDriver.TransactOneRow(sqlStatement, priority, projectName, taskID)
}

func (pd *ProjectData) UpdateTaskSkills(projectName string, taskID int, skills string) error {
	sqlStatement := "UPDATE projects SET skills = $1 WHERE project_name = $2, task_id = $3"
	return pd.PDriver.TransactOneRow(sqlStatement, skills, projectName, taskID)
}
