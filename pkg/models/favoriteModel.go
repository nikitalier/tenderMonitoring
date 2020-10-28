package models

//Favorite ...
type Favorite struct {
	ID       int  `db:"id"`
	TenderID int  `db:"tender_id"`
	UserID   int  `db:"user_id"`
	Status   bool `db:"status"`
}
