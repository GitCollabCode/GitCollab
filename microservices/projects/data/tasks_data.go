package data

import (
	"context"
	"time"
)

type Task struct {
	TaskID          int       `db:"task_id"`
	ProjectID       int       `db:"project_id"`
	ProjectName     string    `db:"project_name"`
	CompletedByID   int       `db:"completed_by_id"`
	CreatedDate     time.Time `db:"date_created_date"`
	CompletedDate   time.Time `db:"completed_date"`
	TaskTitle       string    `db:"task_title"`
	TaskDescription string    `db:"task_description"`
	Diffictly       int       `db:"diffictly"`
	Priority        int       `db:"priority"`
	Skills          []int     `db:"skills"`
}

func (pd *ProjectData) AddTask(taskID int, projectID int, projectName string, createdDate time.Time, taskDescription string, diffictly int, priority int, skills []int) error {
	sqlString :=
		"INSERT INTO tasks(task_id, project_id, project_name, date_created_date, task_description, diffictly, priority, skills)" +
			"VALUES($1, $2, $3, $4, $5, $6, $7, $8)"

	_, err := pd.PDriver.Pool.Exec(context.Background(), sqlString, taskID, projectID, projectName, createdDate, taskDescription, diffictly, priority, skills)
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

func (pd *ProjectData) CompleteTask(taskID int, completedByID string) error {
	sqlStatement := "UPDATE tasks SET completed_by_id = $1, completed_date = $2 WHERE task_id = $3"
	return pd.PDriver.TransactOneRow(sqlStatement, completedByID, time.Now(), taskID)
}

func (pd *ProjectData) UpdateTaskDescription(taskID int, description string) error {
	sqlStatement := "UPDATE projects SET task_description = $1 WHERE task_id = $2"
	return pd.PDriver.TransactOneRow(sqlStatement, description, taskID)
}

func (pd *ProjectData) UpdateTaskDiffictly(taskID int, diffictly string) error {
	sqlStatement := "UPDATE projects SET diffictly = $1 WHERE task_id = $2"
	return pd.PDriver.TransactOneRow(sqlStatement, diffictly, taskID)
}

func (pd *ProjectData) UpdateTaskPriority(taskID int, priority string) error {
	sqlStatement := "UPDATE projects SET priority = $1 WHERE task_id = $2"
	return pd.PDriver.TransactOneRow(sqlStatement, priority, taskID)
}

func (pd *ProjectData) UpdateTaskSkills(taskID int, skills string) error {
	sqlStatement := "UPDATE projects SET skills = $1 WHERE task_id = $2"
	return pd.PDriver.TransactOneRow(sqlStatement, skills, taskID)
}
