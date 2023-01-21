package handlers

import (
	"net/http"
	"strconv"
	"time"

	githubAPI "github.com/GitCollabCode/GitCollab/internal/github"
	jsonio "github.com/GitCollabCode/GitCollab/internal/jsonhttp"
	"github.com/GitCollabCode/GitCollab/internal/models"
	projectModels "github.com/GitCollabCode/GitCollab/microservices/projects/models"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

// GetTasks sends all Tasks under a Project.
func (p *Projects) GetTasks(w http.ResponseWriter, r *http.Request) {
	projectName := chi.URLParam(r, "project-name")

	tasks, err := p.ProjectData.GetTasks(projectName)
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "project has no task"}, w)
		if err != nil {
			p.Log.Fatalf("GetTasks failed to send error response: %s", err)
		}
		return
	}

	if err != nil {
		p.Log.Errorf("GetTasks database search failed: %s", err.Error())
		// NOTE: Repetative code, clean this up
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.Log.Fatalf("GetTasks failed to send error response: %s", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	var res []projectModels.TaskResp
	for _, task := range tasks {
		res = append(res, projectModels.TaskResp{
			TaskID:          task.TaskID,
			ProjectID:       task.ProjectID,
			ProjectName:     task.ProjectName,
			CompletedByID:   task.CompletedByID,
			CreatedDate:     task.CreatedDate,
			CompletedDate:   task.CompletedDate,
			TaskTitle:       task.TaskTitle,
			TaskDescription: task.TaskDescription,
			Diffictly:       task.Diffictly,
			Priority:        task.Priority,
			Skills:          task.Skills,
		})
	}

	err = jsonio.ToJSON(res, w)
	if err != nil {
		p.Log.Fatalf("GetTasks failed to send response: %s", err)
	}
}

// GetTask return a select task.
func (p *Projects) GetTask(w http.ResponseWriter, r *http.Request) {
	projectName := chi.URLParam(r, "project-name")
	taskIDStr := chi.URLParam(r, "task-id")

	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "invalid task ID"}, w)
		if err != nil {
			p.Log.Fatalf("GetTask failed to send error response: %s", err)
		}
		return
	}

	task, err := p.ProjectData.GetTask(projectName, taskID)
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "project has no tasks"}, w)
		if err != nil {
			p.Log.Fatalf("GetTask failed to send error response: %s", err)
		}
		return
	}

	if err != nil {
		p.Log.Errorf("GetTask database search failed: %s", err.Error())
		// NOTE: Repetative code, clean this up
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.Log.Fatalf("GetTask failed to send error response: %s", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	res := projectModels.TaskResp{
		TaskID:          task.TaskID,
		ProjectID:       task.ProjectID,
		ProjectName:     task.ProjectName,
		CompletedByID:   task.CompletedByID,
		CreatedDate:     task.CreatedDate,
		CompletedDate:   task.CompletedDate,
		TaskTitle:       task.TaskTitle,
		TaskDescription: task.TaskDescription,
		Diffictly:       task.Diffictly,
		Priority:        task.Priority,
		Skills:          task.Skills,
	}

	err = jsonio.ToJSON(res, w)
	if err != nil {
		p.Log.Fatalf("GetTask failed to send response: %s", err)
	}
}

// CreateTask creates a task entry in the database.
func (p *Projects) CreateTask(w http.ResponseWriter, r *http.Request) {
	projectName := chi.URLParam(r, "project-name")

	var req projectModels.CreateTaskReq
	err := p.validate.GetJSON(&req, w, r, p.Log)
	if err != nil {
		p.Log.Errorf("CreateTask failed to decode and validate JSON")
		return
	}

	//TODO: Add project owner/admin check

	//TODO: Project ID corresponds to Project name check

	if projectName != req.ProjectName {
		w.WriteHeader(http.StatusBadRequest)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "request body project name and project url miss match"}, w)
		if err != nil {
			p.Log.Fatalf("CreateTask failed to send error response: %s", err)
		}
		return
	}

	err = p.ProjectData.AddTask(
		req.TaskID,
		req.ProjectID,
		req.ProjectName,
		time.Now(),
		req.TaskDescription,
		req.Diffictly,
		req.Priority,
		req.Skills,
	)
	if err != nil {
		p.Log.Errorf("CreateTask database add failed: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.Log.Fatalf("CreateTask failed to send error response: %s", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	err = jsonio.ToJSON(&models.Message{Message: "task created"}, w)
	if err != nil {
		p.Log.Fatalf("CreateTask failed to send success response: %s", err)
	}
}

// DeleteTask delete a select task.
func (p *Projects) DeleteTask(w http.ResponseWriter, r *http.Request) {
	projectName := chi.URLParam(r, "project-name")

	var req projectModels.DeleteTaskReq
	err := p.validate.GetJSON(&req, w, r, p.Log)
	if err != nil {
		p.Log.Errorf("DeleteTask failed to decode and validate JSON")
		return
	}

	//TODO: Add project owner/admin check

	//TODO: add proper error return if task id not found
	err = p.ProjectData.DeleteTask(projectName, req.TaskID)
	if err != nil {
		p.Log.Errorf("DeleteTask database delete failed: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.Log.Fatalf("DeleteTask failed to send error response: %s", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	err = jsonio.ToJSON(&models.Message{Message: "task deleted"}, w)
	if err != nil {
		p.Log.Fatalf("DeleteTask failed to send success response: %s", err)
	}
}

// EditTask edit the details of select task.
func (p *Projects) EditTask(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

// GetRepoIssues retrieve list of issues inside of a repo
func (p *Projects) GetRepoIssues(w http.ResponseWriter, r *http.Request) {
	client, err := githubAPI.GetGitClientFromContext(r)
	if client == nil {
		p.Log.Warning("GetRepoIssues client fetch from context returned nothing!")
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.Log.Fatalf("GetRepoIssues failed to send error response: %s", err)
		}
		return
	}

	if err != nil {
		p.Log.Errorf("GetRepoIssues client fetch from context failed: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.Log.Fatalf("GetRepoIssues failed to send error response: %s", err)
		}
		return
	}

	var repoReq projectModels.RepoInfoReq

	err = p.validate.GetJSON(&repoReq, w, r, p.Log)
	if err != nil {
		p.Log.Errorf("GetRepoIssues failed to decode and validate JSON")
		return
	}

	repo, err := githubAPI.GetRepoByName(client, repoReq.RepoName)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "failed to fetch repo info from github"}, w)
		if err != nil {
			p.Log.Fatalf("GetRepoIssues failed to send error response: %s", err)
		}
		return
	}

	issues, err := githubAPI.GetRepoIssues(client, repo)
	if err != nil {
		p.Log.Error("GetRepoIssues unable to fetch repo issues")
		w.WriteHeader(http.StatusInternalServerError)
		err = jsonio.ToJSON(&models.ErrorMessage{Message: "internal server error"}, w)
		if err != nil {
			p.Log.Fatalf("GetRepoIssues failed to send error response: %s", err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	var issueList []projectModels.RepoIssue
	for _, issue := range issues {
		i := projectModels.RepoIssue{Title: *issue.Title, Url: *issue.URL, State: *issue.State}
		issueList = append(issueList, i)
	}

	resp := projectModels.RepoIssueResp{Issues: issueList}

	err = jsonio.ToJSON(&resp, w)
	if err != nil {
		p.Log.Fatalf("GetRepoIssues failed to send response: %s", err)
	}
}
