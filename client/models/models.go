package models

type Gitevent struct {
	Uuid          string
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
