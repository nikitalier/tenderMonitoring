package repository

import (
	"database/sql"
	"log"

	"github.com/nikitalier/tenderMonitoring/pkg/models"
)

//GetAllTenders ...
func (r *Repository) GetAllTenders() (tenders []models.Tender) {
	err := r.db.Select(&tenders, "select t.id,k.word, description, link, t.add_date, organizer, price from \"Tender\" t join \"Keyword\" k	on t.keyword_id = k.id order by t.add_date desc")
	if err != nil {
		log.Println(err)
	}

	return tenders
}

//GetTenderByID ...
func (r *Repository) GetTenderByID(id int) (tender models.Tender) {
	err := r.db.Get(&tender, "select t.id,k.word, description, link, t.add_date, organizer, price from \"Tender\" t	join \"Keyword\" k on t.keyword_id = k.id where t.id = $1", id)
	if err != nil {
		log.Println(err)
	}

	return tender
}

//GetTenderStatusByID ...
func (r *Repository) GetTenderStatusByID(id int) (status models.TenderStatus) {
	err := r.db.Get(&status, `select s.id, u2.full_name as salname, u.full_name as itname, s.it_status, s.sales_status from "Status" s join "User" u on s.it_user_id = u.id join "User" u2 on s.sales_user_id = u2.id where s.tender_id = $1`, id)
	if err != nil {
		log.Println(err)
	}

	return status
}

func (r *Repository) UpdateTenderStatus(status models.TenderStatus) {
	if status.ITStatus {
		_, err := r.db.NamedExec(`update "Status" set it_status = true, it_user_id = :it_user_id where tender_id = :tender_id`, status)
		if err != nil {
			log.Println(err)
		}
	}

	if status.SalesStatus {
		_, err := r.db.NamedExec(`update "Status" set sales_status = true, sales_user_id = :sales_user_id where tender_id = :tender_id`, status)
		if err != nil {
			log.Println(err)
		}
	}
}

func (r *Repository) CreateTenderStatus(tenderID int) {
	_, err := r.db.Exec(`insert into "Status" (tender_id) values ($1)`, tenderID)
	if err != nil {
		log.Println(err)
	}
}

func (r *Repository) StatusExists(tenderID int) bool {
	var exists bool

	err := r.db.QueryRow(`select exists (select * from "Status" s where s.tender_id = $1)`, tenderID).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
	}
	return exists
}
