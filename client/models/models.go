package models

import "time"

type Gitevent struct {
	Timestamp time.Time
	Repo      string
	Event     string
	Payload   string
}
