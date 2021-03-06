package repository

import (
	"github.com/nikitalier/tenderMonitoring/pkg/models"
)

//CountAllTenders ...
func (r *Repository) CountAllTenders() (count int) {
	err := r.db.Get(&count, `SELECT count(*) FROM "Tender";`)
	if err != nil {
		r.logger.Error().Msg(err.Error())
	}

	return count
}

//CountAllKeywords ...
func (r *Repository) CountAllKeywords() (count int) {
	err := r.db.Get(&count, `SELECT count(*) FROM "Keyword";`)
	if err != nil {
		r.logger.Error().Msg(err.Error())
	}

	return count
}

//GetCountTendersByKeywords ...
func (r *Repository) GetCountTendersByKeywords() (tenders []models.CountTendersByKeywords) {
	err := r.db.Select(&tenders, `select foo.word, count(*) as count from( select k.word from "Tender" t join "TenderMonitoring"."Keyword" k on t.keyword_id  = k.id) as foo group by foo.word order by count desc`)
	if err != nil {
		r.logger.Error().Msg(err.Error())
	}

	return tenders
}

//GetBetsTenders ...
func (r *Repository) GetBetsTenders() (tenders []models.Tender) {
	err := r.db.Select(&tenders, `select distinct t.id, k.word, t.description, t.link, t.organizer, t.price, t.add_date from "Tender" t join "Favorites" f on t.id = f.tender_id join "Keyword" k on t.keyword_id = k.id where f.status = true`)
	if err != nil {
		r.logger.Error().Msg(err.Error())
	}

	return tenders
}

//GetApprovedTenders ...
func (r *Repository) GetApprovedTenders() (tenders []models.Tender) {
	err := r.db.Select(&tenders, `select t.id, k.word, t.description, t.link, t.organizer, t.price, t.add_date from "Tender" t join "Keyword" k on t.keyword_id = k.id join "Status" s on t.id = s.tender_id where s.it_status = true and s.sales_status = true`)
	if err != nil {
		r.logger.Error().Msg(err.Error())
	}

	return tenders
}
