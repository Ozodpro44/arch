package models

import "time"

type Users struct {
	User_id   int
	Username  string
	Gmail     string
	Date      time.Time
	Create_at time.Time
}
