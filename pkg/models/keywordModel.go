package models

import "time"

//Keyword ...
type Keyword struct {
	ID            int       `db:"id"`
	Login         string    `db:"login"`
	UserID        int       `db:"user_id" json:"userID"`
	Word          string    `db:"word" json:"keyword"`
	AddDateTime   time.Time `db:"add_date"`
	AddDateString string    `json:"addDateString"`
}
