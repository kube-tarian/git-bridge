package events

import (
	"github.com/kube-tarian/git-bridge/azure"
	"github.com/kube-tarian/git-bridge/bitbucket"
	"github.com/kube-tarian/git-bridge/gitea"
	"github.com/kube-tarian/git-bridge/github"
	"github.com/kube-tarian/git-bridge/gitlab"
)

var GithubEventTypesSlice = []github.Event{
	"check_run",
	"check_suite",
	"commit_comment",
	"create",
	"delete",
	"deploy_key",
	"deployment",
	"deployment_status",
	"fork",
	"gollum",
	"installation",
	"installation_repositories",
	"integration_installation",
	"integration_installation_repositories",
	"issue_comment",
	"issues",
	"label",
	"member",
	"membership",
	"milestone",
	"meta",
	"organization",
	"org_block",
	"page_build",
	"ping",
	"project_card",
	"project_column",
	"project",
	"public",
	"pull_request",
	"pull_request_review",
	"pull_request_review_comment",
	"push",
	"release",
	"repository",
	"repository_vulnerability_alert",
	"security_advisory",
	"status",
	"team",
	"team_add",
	"watch",
	"workflow_dispatch",
	"workflow_job",
	"workflow_run",
}

var GitlabEventTypesSlice = []gitlab.Event{
	"Push Hook",
	"Tag Push Hook",
	"Issue Hook",
	"Confidential Issue Hook",
	"Note Hook",
	"Merge Request Hook",
	"Wiki Page Hook",
	"Pipeline Hook",
	"Build Hook",
	"Job Hook",
	"System Hook",
}

var BitbucketEventTypesSlice = []bitbucket.Event{
	"repo:push",
	"repo:fork",
	"repo:updated",
	"repo:commit_comment_created",
	"repo:commit_status_created",
	"repo:commit_status_updated",
	"issue:created",
	"issue:updated",
	"issue:comment_created",
	"pullrequest:created",
	"pullrequest:updated",
	"pullrequest:approved",
	"pullrequest:unapproved",
	"pullrequest:fulfilled",
	"pullrequest:rejected",
	"pullrequest:comment_created",
	"pullrequest:comment_updated",
	"pullrequest:comment_deleted",
}

var GiteaEventTypesSlice = []gitea.Event{
	"push",
	"fork",
	"pull_request",
	"issue_comment",
}

var AzureEventTypesSlice = []azure.Event{
	"push",
	"pull",
	"merge",
	"pull_comment",
}
