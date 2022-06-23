package models

import (
	"github.com/google/uuid"
)

type Gitevent struct {
	Uuid          uuid.UUID
	Event         string
	Eventid       string
	Branch        string
	Url           string
	Authorname    string
	Authormail    string
	DoneAt        string
	Repository    string
	Addedfiles    string
	Modifiedfiles string
	Removedfiles  string
	Message       string
}
