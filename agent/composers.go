package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/kube-tarian/git-bridge/azure"
	"github.com/kube-tarian/git-bridge/bitbucket"
	"github.com/kube-tarian/git-bridge/gitea"
	"github.com/kube-tarian/git-bridge/github"
	"github.com/kube-tarian/git-bridge/gitlab"
	"github.com/kube-tarian/git-bridge/models"
)

//gitdatas is an identifier for the gitevent model
//and its used to hold the data from the payloads
var gitdatas models.Gitevent

//gitComposer checks the payload type and extracts the data from the payload and
//compose it into the gitdatas identifier and returns it
func gitComposer(release interface{}, event string) *models.Gitevent {
	uuid := uuid.New().String()

	// here we are using the type assersion. release.(type) will return
	//the type and assingns it to the identifier v
	//the value in v will be used to match with the case
	switch v := release.(type) {

	case github.PushPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Repository.HTMLURL
		gitdatas.Event = event
		gitdatas.Eventid = v.Commits[0].ID
		gitdatas.Authorname = v.Pusher.Name
		gitdatas.Authormail = v.Pusher.Email
		gitdatas.DoneAt = v.HeadCommit.Timestamp
		gitdatas.Repository = v.Repository.Name
		gitdatas.Branch = v.Repository.DefaultBranch
		addedFilesSlice := v.Commits[0].Added
		addedFilesString := getStats(&addedFilesSlice)
		gitdatas.Addedfiles = checkData(addedFilesString)

		modifiedFilesSlice := v.Commits[0].Modified
		modifiedFilesString := getStats(&modifiedFilesSlice)
		gitdatas.Modifiedfiles = checkData(modifiedFilesString)

		removedFilesSlice := v.Commits[0].Removed
		removedFilesString := getStats(&removedFilesSlice)
		gitdatas.Removedfiles = checkData(removedFilesString)

		gitdatas.Message = v.Commits[0].Message

	case github.ForkPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Repository.HTMLURL
		gitdatas.Event = event
		gitdatas.Eventid = strconv.Itoa(int(v.Forkee.ID))
		gitdatas.Authorname = v.Forkee.FullName
		gitdatas.Authormail = "---"
		gitdatas.DoneAt = fmt.Sprintf("%v", v.Forkee.CreatedAt)
		gitdatas.Branch = v.Repository.DefaultBranch
		gitdatas.Addedfiles = "---"
		gitdatas.Modifiedfiles = "---"
		gitdatas.Removedfiles = "---"
		gitdatas.Message = "---"

	case github.PullRequestPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Repository.HTMLURL
		gitdatas.Event = event
		gitdatas.Eventid = strconv.Itoa(int(v.PullRequest.ID))
		gitdatas.Authorname = v.PullRequest.User.Login
		gitdatas.Authormail = "---"
		gitdatas.DoneAt = fmt.Sprintf("%v", v.PullRequest.CreatedAt)
		gitdatas.Repository = v.Repository.Name
		gitdatas.Branch = v.Repository.DefaultBranch
		addedFilesSlice := strconv.Itoa(int(v.PullRequest.Additions))
		addedFilesString := addedFilesSlice
		gitdatas.Addedfiles = checkData(addedFilesString)

		modifiedFilesSlice := strconv.Itoa(int(v.PullRequest.ChangedFiles))
		modifiedFilesString := modifiedFilesSlice
		gitdatas.Modifiedfiles = checkData(modifiedFilesString)

		removedFilesSlice := strconv.Itoa(int(v.PullRequest.Deletions))
		removedFilesString := removedFilesSlice
		gitdatas.Removedfiles = checkData(removedFilesString)

		gitdatas.Message = v.PullRequest.Title

	case gitlab.PushEventPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Project.WebURL
		gitdatas.Event = event
		gitdatas.Eventid = v.Commits[0].ID
		gitdatas.Authorname = v.Commits[0].Author.Name
		gitdatas.Authormail = v.Commits[0].Author.Email
		gitdatas.DoneAt = fmt.Sprintf("%v", v.Commits[0].Timestamp)
		gitdatas.Repository = v.Repository.Name
		gitdatas.Branch = v.Project.DefaultBranch
		addedFilesSlice := v.Commits[0].Added
		addedFilesString := getStats(&addedFilesSlice)
		gitdatas.Addedfiles = checkData(addedFilesString)

		modifiedFilesSlice := v.Commits[0].Modified
		modifiedFilesString := getStats(&modifiedFilesSlice)
		gitdatas.Modifiedfiles = checkData(modifiedFilesString)

		removedFilesSlice := v.Commits[0].Removed
		removedFilesString := getStats(&removedFilesSlice)
		gitdatas.Removedfiles = checkData(removedFilesString)

		gitdatas.Message = v.Commits[0].Message

	case gitlab.MergeRequestEventPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Project.URL
		gitdatas.Event = event
		gitdatas.Eventid = strconv.Itoa(int(v.ObjectAttributes.ID))
		gitdatas.Authorname = v.ObjectAttributes.LastCommit.Author.Name
		gitdatas.Authormail = v.ObjectAttributes.LastCommit.Author.Email
		gitdatas.DoneAt = fmt.Sprintf("%v", v.ObjectAttributes.CreatedAt)
		gitdatas.Repository = v.Repository.Name
		gitdatas.Branch = v.Project.DefaultBranch
		addedFilesSlice := ""
		addedFilesString := addedFilesSlice
		gitdatas.Addedfiles = checkData(addedFilesString)

		modifiedFilesSlice := ""
		modifiedFilesString := modifiedFilesSlice
		gitdatas.Modifiedfiles = checkData(modifiedFilesString)

		removedFilesSlice := ""
		removedFilesString := removedFilesSlice
		gitdatas.Removedfiles = checkData(removedFilesString)

		gitdatas.Message = v.ObjectAttributes.LastCommit.Message

	case bitbucket.RepoPushPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Push.Changes[0].New.Links.HTML.Href
		gitdatas.Event = event
		gitdatas.Eventid = v.Push.Changes[0].New.Target.Hash
		gitdatas.Authorname = v.Push.Changes[0].New.Target.Author.DisplayName
		gitdatas.Authormail = "---"
		gitdatas.DoneAt = fmt.Sprintf("%v", v.Push.Changes[0].New.Target.Date)
		gitdatas.Repository = v.Repository.Name
		gitdatas.Branch = v.Push.Changes[0].New.Name
		addedFilesSlice := ""
		addedFilesString := addedFilesSlice
		gitdatas.Addedfiles = checkData(addedFilesString)

		modifiedFilesSlice := ""
		modifiedFilesString := modifiedFilesSlice
		gitdatas.Modifiedfiles = checkData(modifiedFilesString)

		removedFilesSlice := ""
		removedFilesString := removedFilesSlice
		gitdatas.Removedfiles = checkData(removedFilesString)

		gitdatas.Message = v.Push.Changes[0].New.Target.Message

	case bitbucket.RepoForkPayload:
		rs := time.Now().UTC()
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Fork.Links.HTML.Href
		gitdatas.Event = event
		gitdatas.Eventid = v.Repository.UUID
		gitdatas.Authorname = v.Fork.Owner.DisplayName
		gitdatas.Authormail = "---"
		gitdatas.DoneAt = fmt.Sprintf("%v", rs)
		gitdatas.Repository = v.Repository.Name
		gitdatas.Branch = "---"
		addedFilesSlice := ""
		addedFilesString := addedFilesSlice
		gitdatas.Addedfiles = checkData(addedFilesString)

		modifiedFilesSlice := ""
		modifiedFilesString := modifiedFilesSlice
		gitdatas.Modifiedfiles = checkData(modifiedFilesString)

		removedFilesSlice := ""
		removedFilesString := removedFilesSlice
		gitdatas.Removedfiles = checkData(removedFilesString)

		gitdatas.Message = "---"

	case bitbucket.PullRequestCreatedPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Repository.Links.HTML.Href
		gitdatas.Event = event
		gitdatas.Eventid = v.Repository.UUID
		gitdatas.Authorname = v.PullRequest.Author.DisplayName
		gitdatas.Authormail = "---"
		gitdatas.DoneAt = fmt.Sprintf("%v", v.PullRequest.CreatedOn)
		gitdatas.Repository = v.Repository.Name
		gitdatas.Branch = "---"
		addedFilesSlice := ""
		addedFilesString := addedFilesSlice
		gitdatas.Addedfiles = checkData(addedFilesString)
		modifiedFilesSlice := ""
		modifiedFilesString := modifiedFilesSlice
		gitdatas.Modifiedfiles = checkData(modifiedFilesString)

		removedFilesSlice := ""
		removedFilesString := removedFilesSlice
		gitdatas.Removedfiles = checkData(removedFilesString)

		gitdatas.Message = v.PullRequest.Description

	case azure.PushPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Resource.Repository.RemoteURL
		gitdatas.Event = v.EventType
		gitdatas.Eventid = v.ID
		gitdatas.Authorname = v.Resource.Commits[0].Author.Name
		gitdatas.Authormail = v.Resource.Commits[0].Author.Email
		gitdatas.DoneAt = v.Resource.Commits[0].Author.Date.String()
		gitdatas.Repository = v.Resource.Repository.Name
		gitdatas.Branch = v.Resource.Repository.DefaultBranch
		addedFilesSlice := ""
		addedFilesString := addedFilesSlice
		gitdatas.Addedfiles = checkData(addedFilesString)

		modifiedFilesSlice := ""
		modifiedFilesString := modifiedFilesSlice
		gitdatas.Modifiedfiles = checkData(modifiedFilesString)

		removedFilesSlice := ""
		removedFilesString := removedFilesSlice
		gitdatas.Removedfiles = checkData(removedFilesString)

		gitdatas.Message = v.Resource.Commits[0].Comment

	case azure.PullRequestCreatedPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Resource.Repository.RemoteURL
		gitdatas.Event = v.EventType
		gitdatas.Eventid = v.ID
		gitdatas.Authorname = v.Message.Text
		gitdatas.Authormail = v.Resource.Repository.URL
		gitdatas.DoneAt = v.Resource.CreationDate.String()
		gitdatas.Repository = v.Resource.Repository.Name
		gitdatas.Branch = v.Resource.Repository.DefaultBranch
		addedFilesSlice := ""
		addedFilesString := addedFilesSlice
		gitdatas.Addedfiles = checkData(addedFilesString)

		modifiedFilesSlice := ""
		modifiedFilesString := modifiedFilesSlice
		gitdatas.Modifiedfiles = checkData(modifiedFilesString)

		removedFilesSlice := ""
		removedFilesString := removedFilesSlice
		gitdatas.Removedfiles = checkData(removedFilesString)

		gitdatas.Message = v.Message.Text

	case azure.PullRequestCommentedOnPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Resource.Links.Repository.Href
		gitdatas.Event = v.EventType
		gitdatas.Eventid = v.ID
		gitdatas.Authorname = v.Resource.Author.DisplayName
		gitdatas.Authormail = v.Resource.Author.UniqueName
		gitdatas.DoneAt = v.Resource.LastUpdatedDate.String()
		gitdatas.Repository = "---"
		gitdatas.Branch = "---"
		addedFilesSlice := ""
		addedFilesString := addedFilesSlice
		gitdatas.Addedfiles = checkData(addedFilesString)

		modifiedFilesSlice := ""
		modifiedFilesString := modifiedFilesSlice
		gitdatas.Modifiedfiles = checkData(modifiedFilesString)

		removedFilesSlice := ""
		removedFilesString := removedFilesSlice
		gitdatas.Removedfiles = checkData(removedFilesString)

		gitdatas.Message = fmt.Sprintf("Comment : %v", v.Resource.Content)

	case azure.PullRequestMergeAttemptedPayload:

		gitdatas.Uuid = uuid
		gitdatas.Url = v.Resource.Repository.RemoteURL
		gitdatas.Event = v.EventType
		gitdatas.Eventid = v.ID
		gitdatas.Authorname = v.Message.Text
		gitdatas.Authormail = v.Resource.Repository.URL
		gitdatas.DoneAt = v.Resource.CreationDate.String()
		gitdatas.Repository = v.Resource.Repository.Name
		gitdatas.Branch = v.Resource.Repository.DefaultBranch
		addedFilesSlice := ""
		addedFilesString := addedFilesSlice
		gitdatas.Addedfiles = checkData(addedFilesString)

		modifiedFilesSlice := ""
		modifiedFilesString := modifiedFilesSlice
		gitdatas.Modifiedfiles = checkData(modifiedFilesString)

		removedFilesSlice := ""
		removedFilesString := removedFilesSlice
		gitdatas.Removedfiles = checkData(removedFilesString)

		gitdatas.Message = v.Message.Text

	case gitea.PushEventPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Repository.CloneUrl
		gitdatas.Event = event
		gitdatas.Eventid = v.Repository.Name
		gitdatas.Authorname = v.Sender.Username
		gitdatas.Authormail = v.Sender.Email
		gitdatas.DoneAt = v.Sender.Created.String()
		gitdatas.Repository = v.Repository.Name
		gitdatas.Branch = v.Repository.DefaultBranch
		addedFilesSlice := v.Commits[0].Added
		addedFilesString := getStats(&addedFilesSlice)
		gitdatas.Addedfiles = checkData(addedFilesString)

		modifiedFilesSlice := v.Commits[0].Modified
		modifiedFilesString := getStats(&modifiedFilesSlice)
		gitdatas.Modifiedfiles = checkData(modifiedFilesString)

		removedFilesSlice := v.Commits[0].Removed
		removedFilesString := getStats(&removedFilesSlice)
		gitdatas.Removedfiles = checkData(removedFilesString)

		gitdatas.Message = v.HeadCommit.Message

	case gitea.PullRequestPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Repository.CloneUrl
		gitdatas.Event = event
		gitdatas.Eventid = v.Repository.Name
		gitdatas.Authorname = v.Sender.Username
		gitdatas.Authormail = v.Sender.Email
		gitdatas.DoneAt = v.Sender.Created.String()
		gitdatas.Repository = v.Repository.Name
		gitdatas.Branch = v.Repository.DefaultBranch
		addedFilesSlice := ""
		addedFilesString := addedFilesSlice
		gitdatas.Addedfiles = checkData(addedFilesString)

		modifiedFilesSlice := ""
		modifiedFilesString := modifiedFilesSlice
		gitdatas.Modifiedfiles = checkData(modifiedFilesString)

		removedFilesSlice := ""
		removedFilesString := removedFilesSlice
		gitdatas.Removedfiles = checkData(removedFilesString)

		gitdatas.Message = fmt.Sprintf("Pull Request : %v", v.Action)

	case gitea.PullRequestCommentedPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Repository.CloneUrl
		gitdatas.Event = event
		gitdatas.Eventid = v.Repository.Name
		gitdatas.Authorname = v.Sender.Username
		gitdatas.Authormail = v.Sender.Email
		gitdatas.DoneAt = v.Sender.Created.String()
		gitdatas.Repository = v.Repository.Name
		gitdatas.Branch = v.Repository.DefaultBranch
		addedFilesSlice := ""
		addedFilesString := addedFilesSlice
		gitdatas.Addedfiles = checkData(addedFilesString)

		modifiedFilesSlice := ""

		modifiedFilesString := modifiedFilesSlice
		gitdatas.Modifiedfiles = checkData(modifiedFilesString)

		removedFilesSlice := ""
		removedFilesString := removedFilesSlice
		gitdatas.Removedfiles = checkData(removedFilesString)

		gitdatas.Message = fmt.Sprintf("Comment : %v", v.Comment.Body)

	case gitea.ForkEventPayload:
		gitdatas.Uuid = uuid
		gitdatas.Url = v.Forkee.CloneUrl
		gitdatas.Event = event
		gitdatas.Eventid = v.Forkee.Name
		gitdatas.Authorname = v.Sender.Username
		gitdatas.Authormail = v.Sender.Email
		gitdatas.DoneAt = v.Sender.Created.String()
		gitdatas.Repository = v.Forkee.Name
		gitdatas.Branch = v.Forkee.DefaultBranch
		addedFilesSlice := ""
		addedFilesString := addedFilesSlice
		gitdatas.Addedfiles = checkData(addedFilesString)

		modifiedFilesSlice := ""
		modifiedFilesString := modifiedFilesSlice
		gitdatas.Modifiedfiles = checkData(modifiedFilesString)

		removedFilesSlice := ""
		removedFilesString := removedFilesSlice
		gitdatas.Removedfiles = checkData(removedFilesString)

		gitdatas.Message = v.Forkee.Description

	}
	return &gitdatas
}

//getStats builds a string from the given slice
func getStats(stat *[]string) string {
	var sb strings.Builder
	for _, comm := range *stat {
		sb.WriteString(comm)
		sb.WriteString(",")
	}
	return sb.String()
}

//checkData checks whether the data is empty or not and
//returns "---" if data is empty
func checkData(data string) string {
	if data == "" {
		return "---"
	} else {
		return data
	}
}
