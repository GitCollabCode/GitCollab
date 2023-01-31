package projectWebhooks

import (
	"regexp"
	"strconv"

	"github.com/GitCollabCode/GitCollab/microservices/projects/data"
	"github.com/GitCollabCode/GitCollab/microservices/projects/handlers"
	"github.com/google/go-github/github"
)

type ProjectWebhooks struct {
	pd *data.ProjectData
}

func NewProjectWebhooks(pd *data.ProjectData) *ProjectWebhooks {
	return &ProjectWebhooks{pd}
}

func (pw *ProjectWebhooks) ProjectWebhookHandlers(event interface{}) error {
	var err error = nil
	switch event := event.(type) {
	case *github.PullRequestEvent:
		err = pw.processPullRequest(event)
	case *github.PullRequestReviewEvent:
		err = pw.processPullRequestReview(event)
	case *github.IssueEvent:
		err = pw.processIssues(event)
	}
	return err
}

func (pw *ProjectWebhooks) processPullRequest(event interface{}) error {
	req := event.(*github.PullRequestEvent)

	issueID, err := parseIssueIDFromURL(*req.PullRequest.IssueURL)
	if err != nil {
		return err
	}

	task, err := pw.pd.GetTask(*req.Repo.Name, issueID)
	if err != nil {
		return err
	}

	if task == nil {
		return nil
	}

	if *req.PullRequest.Merged {
		pw.pd.CompleteTask(*req.Repo.Name, issueID, int(*req.PullRequest.Assignee.ID))
		pw.pd.UpdateTaskStatus(*req.Repo.Name, issueID, handlers.TaskStatusCompleted)
		return nil
	}

	return nil
}

func (pw *ProjectWebhooks) processPullRequestReview(event interface{}) error {
	req := event.(*github.PullRequestReviewEvent)

	issueID, err := parseIssueIDFromURL(*req.PullRequest.IssueURL)
	if err != nil {
		return err
	}

	task, err := pw.pd.GetTask(*req.Repo.Name, issueID)
	if err != nil {
		return err
	}

	if task == nil {
		return nil
	}

	if *req.Review.State == "APPROVED" {
		err = pw.pd.UpdateTaskStatus(*req.Repo.Name, issueID, handlers.TaskStatusApproved)
	} else if *req.Review.State == "CHANGES_REQUESTED" {
		err = pw.pd.UpdateTaskStatus(*req.Repo.Name, issueID, handlers.TaskStatusChangesRequested)
	} else if *req.Review.State == "DISMISSED" {
		err = pw.pd.UpdateTaskStatus(*req.Repo.Name, issueID, handlers.TaskStatusDismissed)
	}

	return err
}

func (pw *ProjectWebhooks) processIssues(event interface{}) error {
	return nil
}

func parseIssueIDFromURL(url string) (int, error) {
	r, err := regexp.Compile(`https://github\\.com/\\w+/\\w+/issues/(?P<issue_id>[0-9]+)`)
	if err != nil {
		return -1, err
	}

	matches := r.FindStringSubmatch(url)
	issueIDIndex := r.SubexpIndex("issue_id")

	issueID, err := strconv.Atoi(matches[issueIDIndex])
	if err != nil {
		return -1, err
	}

	return issueID, nil
}
