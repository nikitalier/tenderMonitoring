package models

import (
	"time"
)

//Tender ...
type Tender struct {
	ID          int       `db:"id"`
	Keyword     string    `db:"word"`
	Description string    `db:"description"`
	Link        string    `db:"link"`
	AddDate     time.Time `db:"add_date"`
	Organizer   string    `db:"organizer"`
	Price       string    `db:"price"`
}

//TenderStatus ...
type TenderStatus struct {
	ID            int    `db:"id"`
	TenderID      int    `db:"tender_id"`
	ITUserName    string `db:"itname"`
	SalesUserName string `db:"salname"`
	SalesStatus   bool   `db:"sales_status"`
	ITStatus      bool   `db:"it_status"`
	ITUserID      int    `db:"it_user_id"`
	SalesUserID   int    `db:"sales_user_id"`
}
