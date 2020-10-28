package repository

import (
	"database/sql"
	"log"

	"github.com/nikitalier/tenderMonitoring/pkg/models"
)

//FavoriteExists ...
func (r *Repository) FavoriteExists(tenderID int, userID int) bool {
	var exists bool

	err := r.db.QueryRow("select exists (select * from \"Favorites\" f where f.tender_id = $1 and f.user_id = $2)", tenderID, userID).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
	}
	return exists
}

//AddNewFavorite ...
func (r *Repository) AddNewFavorite(fav models.Favorite) {
	// sql := "insert into \"Favorites\" (user_id, tender_id , status ) values ($1, $2, $3)"
	_, err := r.db.Exec("insert into \"Favorites\" (user_id, tender_id , status ) values ($1, $2, $3)", fav.UserID, fav.TenderID, fav.Status)
	if err != nil {
		log.Println(err)
	}
}

//FindFavorite ...
func (r *Repository) FindFavorite(tednerID int, userID int) (f models.Favorite) {
	err := r.db.Get(&f, "select * from \"Favorites\" f where user_id = $1 and tender_id = $2", userID, tednerID)
	if err != nil {
		log.Println(err)
	}
	// log.Println(f)
	return f
}

//UpdateFavorite ...
func (r *Repository) UpdateFavorite(f models.Favorite) {
	// log.Println(f)
	_, err := r.db.NamedExec(`update "Favorites" set status = :status where user_id = :user_id and tender_id = :tender_id`, f)
	if err != nil {
		log.Println(err)
	}
}
