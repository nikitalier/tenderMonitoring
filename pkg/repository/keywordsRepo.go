package repository

import (
	"github.com/nikitalier/tenderMonitoring/pkg/models"
)

//GetAllKeyWords ...
func (r *Repository) GetAllKeyWords() (keywords []models.Keyword) {
	err := r.db.Select(&keywords, "select k.id, login, word, add_date from \"Keyword\" k join \"User\" u on user_id = u.id")
	if err != nil {
		r.logger.Error().Msg(err.Error())
	}

	return keywords
}

//DeleteKeyWordsByIDs ...
func (r *Repository) DeleteKeyWordsByIDs(ids string) bool {
	_, err := r.db.Exec("delete from \"Keyword\" where id in (" + ids + ")")
	if err != nil {
		r.logger.Error().Msg(err.Error())
		return false
	}
	return true
}

//AddKeyword ...
func (r *Repository) AddKeyword(keyword models.Keyword) (id int, result bool) {
	rows, err := r.db.NamedQuery("insert into \"Keyword\" (user_id, word, add_date) values (:user_id, :word, :add_date) RETURNING id", keyword)
	if err != nil {
		r.logger.Error().Msg(err.Error())
		return 0, false
	}

	if rows.Next() {
		rows.Scan(&id)
	}

	return id, true
}
