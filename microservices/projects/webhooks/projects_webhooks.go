package projectWebhooks

import "github.com/google/go-github/github"

func ProjectWebhookHandlers(event interface{}) error {
	switch event := event.(type) {
	case *github.CommitCommentEvent:
		processCommitCommentEvent(event)
	case *github.CreateEvent:
		processCreateEvent(event)
	}
}
