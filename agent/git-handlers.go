package main

import (
	"log"

	"net/http"

	"github.com/kube-tarian/git-bridge/azure"
	"github.com/kube-tarian/git-bridge/bitbucket"
	"github.com/kube-tarian/git-bridge/github"
	"github.com/kube-tarian/git-bridge/gitlab"
)

func (app *application) azureHandler(w http.ResponseWriter, r *http.Request) {
	event := r.Header.Get("X-Azure-Event")
	if event == "" {
		log.Println("ErrMissingGithubEventHeader")
	}
	hook, _ := azure.New()
	payload, err := hook.Parse(r, azure.PushPayloadEvent, azure.PullRequestCommentEvent, azure.PullRequestCreatedEvent, azure.PullRequestMergeAttemptedEvent)
	if err != nil {
		if err == azure.ErrEventNotFound {
			log.Print("Error This Event is not Supported")
		}
	}
	switch value := payload.(type) {
	case azure.PushPayload:
		release := value
		composed := gitComposer(release, event)
		app.publish.JS.GitPublish(composed)
	case azure.PullRequestCreatedPayload:
		release := value
		composed := gitComposer(release, event)
		app.publish.JS.GitPublish(composed)
	case azure.PullRequestCommentedOnPayload:
		release := value
		composed := gitComposer(release, event)
		app.publish.JS.GitPublish(composed)
	case azure.PullRequestMergeAttemptedPayload:
		release := value
		composed := gitComposer(release, event)
		app.publish.JS.GitPublish(composed)
	}
}

//githubHandler handles the github webhooks post requests.
func (app *application) githubHandler(w http.ResponseWriter, r *http.Request) {
	event := r.Header.Get("X-GitHub-Event")
	if event == "" {
		log.Println("ErrMissingGithubEventHeader")
	}
	hook, _ := github.New()
	payload, err := hook.Parse(r, github.PushEvent, github.ForkEvent, github.PullRequestEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			log.Print("Error This Event is not Supported")
		}

	}

	switch value := payload.(type) {
	case github.PushPayload:
		release := value
		composed := gitComposer(release, event)
		app.publish.JS.GitPublish(composed)
	case github.ForkPayload:
		release := value
		composed := gitComposer(release, event)
		app.publish.JS.GitPublish(composed)
	case github.PullRequestPayload:
		release := value
		composed := gitComposer(release, event)
		app.publish.JS.GitPublish(composed)

	}

}

//gitlabHandler handles the github webhooks post requests.
func (app *application) gitlabHandler(w http.ResponseWriter, r *http.Request) {
	event := r.Header.Get("X-Gitlab-Event")
	if len(event) == 0 {
		log.Println("ErrMissingGitLabEventHeader")
	}
	hook, _ := gitlab.New()
	payload, err := hook.Parse(r, gitlab.PushEvents, gitlab.MergeRequestEvents)
	if err != nil {
		if err == gitlab.ErrEventNotFound {
			log.Print("Error This Event is not Supported")
		}
	}

	switch value := payload.(type) {

	case gitlab.PushEventPayload:
		release := value
		composed := gitComposer(release, event)
		app.publish.JS.GitPublish(composed)
	case gitlab.MergeRequest:
		release := value
		composed := gitComposer(release, event)
		app.publish.JS.GitPublish(composed)
	}
}

//bitBucketHandler handles the github webhooks post requests.
func (app *application) bitBucketHandler(w http.ResponseWriter, r *http.Request) {
	event := r.Header.Get("X-Event-Key")
	if event == "" {
		log.Println("ErrMissingEventKeyHeader")
	}
	hook, _ := bitbucket.New()
	payload, err := hook.Parse(r, bitbucket.RepoPushEvent, bitbucket.RepoForkEvent, bitbucket.PullRequestCreatedEvent)
	if err != nil {
		if err == github.ErrEventNotFound {
			log.Print("Error This Event is not Supported")
		}
	}

	switch value := payload.(type) {

	case bitbucket.RepoPushPayload:
		release := value
		composed := gitComposer(release, event)
		app.publish.JS.GitPublish(composed)
	case bitbucket.RepoForkPayload:
		release := value
		composed := gitComposer(release, event)
		app.publish.JS.GitPublish(composed)
	case bitbucket.PullRequestCreatedPayload:
		release := value
		composed := gitComposer(release, event)
		app.publish.JS.GitPublish(composed)

	}
}
