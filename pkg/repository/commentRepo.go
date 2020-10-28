package repository

import (
	"log"

	"github.com/nikitalier/tenderMonitoring/pkg/models"
)

//GetAllCommentsByTenderID ...
func (r *Repository) GetAllCommentsByTenderID(tenderID int) (comments []models.Comment) {
	err := r.db.Select(&comments, `select c.id, u.full_name, c."text" from "Comment" c join "TenderMonitoring"."User" u on c.user_id = u.id where c.tender_id = $1`, tenderID)
	if err != nil {
		log.Println(err)
	}
	return comments
}

//AddNewComment ...
func (r *Repository) AddNewComment(comment models.Comment) {
	_, err := r.db.Exec(`insert into "Comment"  (tender_id, user_id, "text") values ($1, $2, $3)`, comment.TenderID, comment.UserID, comment.Text)
	if err != nil {
		log.Println(err)
	}
}
