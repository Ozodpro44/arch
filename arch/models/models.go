package models

import "time"

type Users struct {
	User_id 	int
	User_name	string
	Gmail		string
	Date		time.Time
	Create_at 	time.Time
}