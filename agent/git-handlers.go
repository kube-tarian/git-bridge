package main

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kube-tarian/git-bridge/azure"
	"github.com/kube-tarian/git-bridge/bitbucket"
	"github.com/kube-tarian/git-bridge/events"
	"github.com/kube-tarian/git-bridge/gitea"
	"github.com/kube-tarian/git-bridge/github"
	"github.com/kube-tarian/git-bridge/gitlab"
)

func (app *application) giteaHandler(c *gin.Context) {
	repo := "Gitea"
	event := c.Request.Header.Get("X-Gitea-Event")
	if event == "" {
		log.Println("ErrMissingGiteaEventHeader")
	}
	hook, _ := gitea.New()
	payload, err := hook.Parse(c.Request, events.GiteaEventTypesSlice...)
	if err != nil {
		if err == gitea.ErrEventNotFound {
			log.Print("Error This Event is not Supported")
		}
	}

	switch value := payload.(type) {

	case gitea.PushEventPayload:
		release := value
		app.publish.JS.Publish(repo, "PushEventPayload", JtoS(&release))
	case gitea.PullRequestPayload:
		release := value
		app.publish.JS.Publish(repo, "PullRequestPayload", JtoS(&release))

	case gitea.PullRequestCommentedPayload:
		release := value
		app.publish.JS.Publish(repo, "PullRequestCommentedPayload", JtoS(&release))
	case gitea.ForkEventPayload:
		release := value
		app.publish.JS.Publish(repo, "ForkEventPayload", JtoS(&release))
}

func (app *application) azureHandler(c *gin.Context) {
	repo:= "Azure"
	event := c.Request.Header.Get("X-Azure-Event")
	if event == "" {
		log.Println("ErrMissingGithubEventHeader")
	}
	hook, _ := azure.New()
	payload, err := hook.Parse(c.Request, events.AzureEventTypesSlice...)
	if err != nil {
		if err == azure.ErrEventNotFound {
			log.Print("Error This Event is not Supported")
		}
	}
	switch value := payload.(type) {
	case azure.PushPayload:
		release := value
		app.publish.JS.Publish(repo, "PushPayload", JtoS(&release))
	case azure.PullRequestCreatedPayload:
		release := value
		app.publish.JS.Publish(repo, "PullRequestCreatedPayload", JtoS(&release))
	case azure.PullRequestCommentedOnPayload:
		release := value
		app.publish.JS.Publish(repo, "PullRequestCommentedOnPayload", JtoS(&release))
	case azure.PullRequestMergeAttemptedPayload:
		release := value
		app.publish.JS.Publish(repo, "PullRequestMergeAttemptedPayload", JtoS(&release))

	}
}

//githubHandler handles the github webhooks post requests.
func (app *application) githubHandler(c *gin.Context) {
	repo:= "Github"
	event := c.Request.Header.Get("X-GitHub-Event")
	if event == "" {
		log.Println("ErrMissingGithubEventHeader")
	}
	hook, _ := github.New()
	payload, err := hook.Parse(c.Request, events.GithubEventTypesSlice...)
	if err != nil {
		if err == github.ErrEventNotFound {
			log.Print("Error This Event is not Supported")
		}

	}

	switch value := payload.(type) {
	case github.CheckRunPayload:
		release := value
		app.publish.JS.Publish(repo, "CheckRunPayload", JtoS(&release))
	case github.CheckSuitePayload:
		release := value
		app.publish.JS.Publish(repo, "CheckSuitePayload", JtoS(&release))
	case github.CommitCommentPayload:
		release := value
		app.publish.JS.Publish(repo, "CommitCommentPayload", JtoS(&release))
	case github.CreatePayload:
		release := value
		app.publish.JS.Publish(repo, "CreatePayload", JtoS(&release))
	case github.DeletePayload:
		release := value
		app.publish.JS.Publish(repo, "DeletePayload", JtoS(&release))
	case github.DeploymentPayload:
		release := value
		app.publish.JS.Publish(repo, "DeploymentPayload", JtoS(&release))
	case github.DeploymentStatusPayload:
		release := value
		app.publish.JS.Publish(repo, "DeploymentStatusPayload", JtoS(&release))
	case github.ForkPayload:
		release := value
		app.publish.JS.Publish(repo, "ForkPayload", JtoS(&release))
	case github.GollumPayload:
		release := value
		app.publish.JS.Publish(repo, "GollumPayload", JtoS(&release))
	case github.InstallationPayload:
		release := value
		app.publish.JS.Publish(repo, "InstallationPayload", JtoS(&release))
	case github.InstallationRepositoriesPayload:
		release := value
		app.publish.JS.Publish(repo, "InstallationRepositoriesPayload", JtoS(&release))
	case github.IssueCommentPayload:
		release := value
		app.publish.JS.Publish(repo, "IssueCommentPayload", JtoS(&release))
	case github.IssuesPayload:
		release := value
		app.publish.JS.Publish(repo, "IssuesPayload", JtoS(&release))
	case github.LabelPayload:
		release := value
		app.publish.JS.Publish(repo, "LabelPayload", JtoS(&release))
	case github.MemberPayload:
		release := value
		app.publish.JS.Publish(repo, "MemberPayload", JtoS(&release))
	case github.MembershipPayload:
		release := value
		app.publish.JS.Publish(repo, "MembershipPayload", JtoS(&release))
	case github.MilestonePayload:
		release := value
		app.publish.JS.Publish(repo, "MilestonePayload", JtoS(&release))
	case github.OrganizationPayload:
		release := value
		app.publish.JS.Publish(repo, "OrganizationPayload", JtoS(&release))
	case github.PageBuildPayload:
		release := value
		app.publish.JS.Publish(repo, "PageBuildPayload", JtoS(&release))
	case github.PingPayload:
		release := value
		app.publish.JS.Publish(repo, "PingPayload", JtoS(&release))
	case github.ProjectPayload:
		release := value
		app.publish.JS.Publish(repo, "ProjectPayload", JtoS(&release))
	case github.ProjectCardPayload:
		release := value	
		app.publish.JS.Publish(repo, "ProjectCardPayload", JtoS(&release))
	case github.ProjectColumnPayload:
		release := value
		app.publish.JS.Publish(repo, "ProjectColumnPayload", JtoS(&release))
	case github.PublicPayload:
		release := value
		app.publish.JS.Publish(repo, "PublicPayload", JtoS(&release))
	case github.PullRequestPayload:
		release := value
		app.publish.JS.Publish(repo, "PullRequestPayload", JtoS(&release))
	case github.PullRequestReviewPayload:
		release := value
		app.publish.JS.Publish(repo, "PullRequestReviewPayload", JtoS(&release))
	case github.PullRequestReviewCommentPayload:
		release := value
		app.publish.JS.Publish(repo, "PullRequestReviewCommentPayload", JtoS(&release))
	case github.PushPayload:
		release := value
		app.publish.JS.Publish(repo, "PushPayload", JtoS(&release))
	case github.ReleasePayload:
		release := value
		app.publish.JS.Publish(repo, "ReleasePayload", JtoS(&release))
	case github.RepositoryPayload:
		release := value
		app.publish.JS.Publish(repo, "RepositoryPayload", JtoS(&release))
	case github.StatusPayload:
		release := value
		app.publish.JS.Publish(repo, "StatusPayload", JtoS(&release))
	case github.TeamPayload:
		release := value
		app.publish.JS.Publish(repo, "TeamPayload", JtoS(&release))
	case github.TeamAddPayload:
		release := value
		app.publish.JS.Publish(repo, "TeamAddPayload", JtoS(&release))
	case github.WatchPayload:
		release := value
		app.publish.JS.Publish(repo, "WatchPayload", JtoS(&release))
		case github.WorkflowDispatchPayload:
		release := value
		app.publish.JS.Publish(repo, "WorkflowDispatchPayload", JtoS(&release))
	case github.WorkflowRunPayload:
		release := value
		app.publish.JS.Publish(repo, "WorkflowRunPayload", JtoS(&release))
	case github.WorkflowJobPayload:
		release := value
		app.publish.JS.Publish(repo, "WorkflowJobPayload", JtoS(&release))
	}

}

//gitlabHandler handles the github webhooks post requests.
func (app *application) gitlabHandler(c *gin.Context) {
	repo := "Gitlab"
	event := c.Request.Header.Get("X-Gitlab-Event")
	if len(event) == 0 {
		log.Println("ErrMissingGitLabEventHeader")
	}
	hook, _ := gitlab.New()
	payload, err := hook.Parse(c.Request, events.GitlabEventTypesSlice...)
	if err != nil {
		if err == gitlab.ErrEventNotFound {
			log.Print("Error This Event is not Supported")
		}
	}

	switch value := payload.(type) {


case gitlab.PushEventPayload:
	release := value
	app.publish.JS.Publish(repo, "PushEventPayload", JtoS(&release))
	case gitlab.TagEventPayload:
		release := value
		app.publish.JS.Publish(repo, "TagEventPayload", JtoS(&release))
	case gitlab.IssueEventPayload:
		release := value
		app.publish.JS.Publish(repo, "IssueEventPayload", JtoS(&release))
	case gitlab.ConfidentialIssueEventPayload:
		release := value
		app.publish.JS.Publish(repo, "ConfidentialIssueEventPayload", JtoS(&release))
	case gitlab.CommentEventPayload:
		release := value
		app.publish.JS.Publish(repo, "CommentEventPayload", JtoS(&release))
	case gitlab.MergeRequestEventPayload:
		release := value
		app.publish.JS.Publish(repo, "MergeRequestEventPayload", JtoS(&release))
	case gitlab.WikiPageEventPayload:
		release := value
		app.publish.JS.Publish(repo, "WikiPageEventPayload", JtoS(&release))
	case gitlab.PipelineEventPayload:
		release := value
		app.publish.JS.Publish(repo, "PipelineEventPayload", JtoS(&release))
	case gitlab.BuildEventPayload:
		release := value
		app.publish.JS.Publish(repo, "BuildEventPayload", JtoS(&release))
	case gitlab.JobEventPayload:
		release := value
		app.publish.JS.Publish(repo, "JobEventPayload", JtoS(&release))
	}
}

//bitBucketHandler handles the github webhooks post requests.
func (app *application) bitBucketHandler(c *gin.Context) {
	repo := "BitBucket"
	event := c.Request.Header.Get("X-Event-Key")
	if event == "" {
		log.Println("ErrMissingEventKeyHeader")
	}
	hook, _ := bitbucket.New()
	payload, err := hook.Parse(c.Request, events.BitbucketEventTypesSlice...)
	if err != nil {
		if err == github.ErrEventNotFound {
			log.Print("Error This Event is not Supported")
		}
	}

	switch value := payload.(type) {

	case bitbucket.RepoPushPayload:
		release := value
		app.publish.JS.Publish(repo, "RepoPushPayload", JtoS(&release))
	case bitbucket.RepoForkPayload:
		release := value
		app.publish.JS.Publish(repo, "RepoForkPayload", JtoS(&release))
	case bitbucket.RepoUpdatedPayload:
		release := value
		app.publish.JS.Publish(repo, "RepoUpdatedPayload", JtoS(&release))
	case bitbucket.RepoCommitCommentCreatedPayload:
		release := value
		app.publish.JS.Publish(repo, "RepoCommitCommentCreatedPayload", JtoS(&release))
	case bitbucket.RepoCommitStatusUpdatedPayload:
		release := value
		app.publish.JS.Publish(repo, "RepoCommitStatusUpdatedPayload", JtoS(&release))
	case bitbucket.IssueCreatedPayload:
		release := value
		app.publish.JS.Publish(repo, "IssueCreatedPayload", JtoS(&release))
	case bitbucket.IssueUpdatedPayload:
		release := value
		app.publish.JS.Publish(repo, "IssueUpdatedPayload", JtoS(&release))
	case bitbucket.IssueCommentCreatedPayload:
		release := value
		app.publish.JS.Publish(repo, "IssueCommentCreatedPayload", JtoS(&release))
	case bitbucket.PullRequestCreatedPayload:
		release := value
		app.publish.JS.Publish(repo, "PullRequestCreatedPayload", JtoS(&release))
		case bitbucket.PullRequestUpdatedPayload:
		release := value
		app.publish.JS.Publish(repo, "PullRequestUpdatedPayload", JtoS(&release))
		case bitbucket.PullRequestApprovedPayload:
		release := value
		app.publish.JS.Publish(repo, "PullRequestApprovedPayload", JtoS(&release))
		case bitbucket.PullRequestUnapprovedPayload:
		release := value
		app.publish.JS.Publish(repo, "PullRequestUnapprovedPayload", JtoS(&release))
		case bitbucket.PullRequestMergedPayload:
		release := value
		app.publish.JS.Publish(repo, "PullRequestMergedPayload", JtoS(&release))
		case bitbucket.PullRequestDeclinedPayload:
		release := value
		app.publish.JS.Publish(repo, "PullRequestDeclinedPayload", JtoS(&release))
		case bitbucket.PullRequestCommentCreatedPayload:
		release := value
		app.publish.JS.Publish(repo, "PullRequestCommentCreatedPayload", JtoS(&release))
		case bitbucket.PullRequestCommentUpdatedPayload:
		release := value
		app.publish.JS.Publish(repo, "PullRequestCommentUpdatedPayload", JtoS(&release))
		case bitbucket.PullRequestCommentDeletedPayload:
		release := value
		app.publish.JS.Publish(repo, "PullRequestCommentDeletedPayload", JtoS(&release))
	}
}

func JtoS(j interface{}) string {
	b, _ := json.Marshal(j)
	return string(b)
}
