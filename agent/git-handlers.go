package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kube-tarian/git-bridge/azure"
	"github.com/kube-tarian/git-bridge/bitbucket"
	"github.com/kube-tarian/git-bridge/events"
	"github.com/kube-tarian/git-bridge/gitea"
	"github.com/kube-tarian/git-bridge/github"
	"github.com/kube-tarian/git-bridge/gitlab"
)

func SpreadGithubEvents(value []string) []github.Event {
	v := make([]github.Event, len(value))
	for _, j := range value {

		v = append(v, github.Event(j))

	}
	return v
}

func (app *application) giteaHandler(c *gin.Context) {
	event := c.Request.Header.Get("X-Gitea-Event")
	if event == "" {
		log.Println("ErrMissingGiteaEventHeader")
	}
	hook, _ := gitea.New()
	payload, err := hook.Parse(c.Request, gitea.PushEvent, gitea.PullRequestEvent, gitea.PullRequestReviewEvent, gitea.ForkEvent)
	if err != nil {
		if err == gitea.ErrEventNotFound {
			log.Print("Error This Event is not Supported")
		}
	}

	switch value := payload.(type) {

	case gitea.PushEventPayload:
		release := value
		app.publish.JS.Publish(&release)
	case gitea.PullRequestPayload:
		release := value
		app.publish.JS.Publish(&release)

	case gitea.PullRequestCommentedPayload:
		release := value
		app.publish.JS.Publish(&release)
	case gitea.ForkEventPayload:
		release := value
		app.publish.JS.Publish(&release)

	}
}

func (app *application) azureHandler(c *gin.Context) {
	event := c.Request.Header.Get("X-Azure-Event")
	if event == "" {
		log.Println("ErrMissingGithubEventHeader")
	}
	hook, _ := azure.New()
	payload, err := hook.Parse(c.Request, azure.PushEvent, azure.PullRequestCommentEvent, azure.PullRequestCreatedEvent, azure.PullRequestMergeAttemptedEvent)
	if err != nil {
		if err == azure.ErrEventNotFound {
			log.Print("Error This Event is not Supported")
		}
	}
	switch value := payload.(type) {
	case azure.PushPayload:
		release := value
		app.publish.JS.Publish(&release)
	case azure.PullRequestCreatedPayload:
		release := value
		app.publish.JS.Publish(&release)
	case azure.PullRequestCommentedOnPayload:
		release := value
		app.publish.JS.Publish(&release)
	case azure.PullRequestMergeAttemptedPayload:
		release := value
		app.publish.JS.Publish(&release)
	}
}

//githubHandler handles the github webhooks post requests.
func (app *application) githubHandler(c *gin.Context) {

	value := SpreadGithubEvents(events.GithubEventTypesSlice)
	event := c.Request.Header.Get("X-GitHub-Event")
	if event == "" {
		log.Println("ErrMissingGithubEventHeader")
	}
	hook, _ := github.New()
	payload, err := hook.Parse(c.Request, value...)
	if err != nil {
		if err == github.ErrEventNotFound {
			log.Print("Error This Event is not Supported")
		}

	}

	switch value := payload.(type) {
	case github.PushPayload:
		release := value
		app.publish.JS.Publish(&release)
	case github.ForkPayload:
		release := value
		app.publish.JS.Publish(&release)
	case github.PullRequestPayload:
		release := value
		app.publish.JS.Publish(&release)

	}

}

//gitlabHandler handles the github webhooks post requests.
func (app *application) gitlabHandler(c *gin.Context) {
	event := c.Request.Header.Get("X-Gitlab-Event")
	if len(event) == 0 {
		log.Println("ErrMissingGitLabEventHeader")
	}
	hook, _ := gitlab.New()
	payload, err := hook.Parse(c.Request, gitlab.PushEvents, gitlab.MergeRequestEvents)
	if err != nil {
		if err == gitlab.ErrEventNotFound {
			log.Print("Error This Event is not Supported")
		}
	}

	switch value := payload.(type) {

	case gitlab.PushEventPayload:
		release := value
		app.publish.JS.Publish(&release)
	case gitlab.MergeRequest:
		release := value
		app.publish.JS.Publish(&release)
	}
}

//bitBucketHandler handles the github webhooks post requests.
func (app *application) bitBucketHandler(c *gin.Context) {
	event := c.Request.Header.Get("X-Event-Key")
	if event == "" {
		log.Println("ErrMissingEventKeyHeader")
	}
	hook, _ := bitbucket.New()
	payload, err := hook.Parse(c.Request, bitbucket.RepoPushEvent, bitbucket.RepoForkEvent, bitbucket.PullRequestCreatedEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			log.Print("Error This Event is not Supported")
		}
	}

	switch value := payload.(type) {

	case bitbucket.RepoPushPayload:
		release := value
		app.publish.JS.Publish(&release)
	case bitbucket.RepoForkPayload:
		release := value
		app.publish.JS.Publish(&release)
	case bitbucket.PullRequestCreatedPayload:
		release := value
		app.publish.JS.Publish(&release)

	}
}
