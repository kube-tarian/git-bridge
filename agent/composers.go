package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/vijeyash1/gitevent/bitbucket"
	"github.com/vijeyash1/gitevent/github"
	"github.com/vijeyash1/gitevent/gitlab"
	"github.com/vijeyash1/gitevent/models"
)

//gitdatas is an identifier for the gitevent model
//and its used to hold the data from the payloads
var gitdatas models.Gitevent

//gitComposer checks the payload type and extracts the data from the payload and
//compose it into the gitdatas identifier and returns it
func gitComposer(release interface{}, event string) *models.Gitevent {
	uuid := uuid.New()

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
