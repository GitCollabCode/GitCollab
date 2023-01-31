package projectsModels

import "time"

// List of repo issues response
// swagger:response repoIssueResp
type _ struct {
	// in:body
	// Required: true
	Body RepoIssueResp
}

type RepoIssueResp struct {
	// Github repositories
	Issues []RepoIssue `json:"issues"`
}

type CreateTaskReq struct {
	TaskID          int    `json:"task_id"`
	ProjectID       int    `json:"project_id"`
	ProjectName     string `json:"project_name"`
	TaskTitle       string `json:"task_title"`
	TaskDescription string `json:"task_description"`
	Diffictly       int    `json:"diffictly"`
	Priority        int    `json:"priority"`
	Skills          []int  `json:"skills"`
}

type TaskResp struct {
	TaskID          int       `json:"task_id"`
	ProjectID       int       `json:"project_id"`
	ProjectName     string    `json:"project_name"`
	CompletedByID   int       `json:"completed_by_id"`
	CreatedDate     time.Time `json:"date_created_date"`
	CompletedDate   time.Time `json:"completed_date"`
	TaskTitle       string    `json:"task_title"`
	TaskDescription string    `json:"task_description"`
	Diffictly       int       `json:"diffictly"`
	Priority        int       `json:"priority"`
	Skills          []int     `json:"skills"`
}

type DeleteTaskReq struct {
	TaskID int `json:"task_id"`
}

type EditTaskReq struct {
	TaskTitle       string `json:"task_title"`
	TaskDescription string `json:"task_description"`
	Diffictly       int    `json:"diffictly"`
	Priority        int    `json:"priority"`
	Skills          []int  `json:"skills"`
}
