package repository

import (
	"log"

	"github.com/nikitalier/tenderMonitoring/pkg/models"
)

//GetAllKeyWords ...
func (r *Repository) GetAllKeyWords() (keywords []models.Keyword) {
	err := r.db.Select(&keywords, "select k.id, login, word, add_date from \"Keyword\" k join \"User\" u on user_id = u.id")
	if err != nil {
		log.Println(err)
	}
	//log.Println(keywords[0].AddDate.Format("02-01-2006"))
	return keywords
}

//DeleteKeyWordsByIDs ...
func (r *Repository) DeleteKeyWordsByIDs(ids string) bool {
	// var idsString string

	// for i := 0; i < len(ids); i++ {
	// 	idsString += fmt.Sprint(ids[i]) + ","
	// }

	// deleteSQL := "delete from \"Keyword\" where id in (" + idsString[:len(idsString)-1] + ")"
	deleteSQL := "delete from \"Keyword\" where id in (" + ids + ")"
	// log.Println(deleteSQL)
	// _, err := r.db.Exec("delete from \"Keyword\" where id in ($1)", idsString[:len(idsString)-1])
	_, err := r.db.Exec(deleteSQL)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

//AddKeyword ...
// func (r *Repository) AddKeyword(keyword models.Keyword) bool {
// 	// sql := "insert into \"Keyword\" (user_id, word, add_date) values ($1, '$2', '$3')"

// 	_, err := r.db.Exec("insert into \"Keyword\" (user_id, word, add_date) values ($1, $2, $3)", keyword.UserID, keyword.Word, keyword.AddDateTime)
// 	if err != nil {
// 		log.Println(err)
// 		return false
// 	}
// 	return true
// }

//AddKeyword ...
func (r *Repository) AddKeyword(keyword models.Keyword) (id int, result bool) {
	rows, err := r.db.NamedQuery("insert into \"Keyword\" (user_id, word, add_date) values (:user_id, :word, :add_date) RETURNING id", keyword)
	if err != nil {
		log.Println(err)
		return 0, false
	}

	// log.Println(rows.Scan())

	if rows.Next() {
		rows.Scan(&id)
	}

	// log.Println(id)

	return id, true
}
