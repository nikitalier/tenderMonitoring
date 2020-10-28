package models

//Comment ...
type Comment struct {
	ID           int    `db:"id"`
	TenderID     int    `db:"tender_id"`
	UserID       int    `db:"user_id"`
	UserFullName string `db:"full_name"`
	Text         string `db:"text"`
}
