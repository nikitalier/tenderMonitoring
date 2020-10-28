package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
)

//New repository - add db in repository
func New(pr *sqlx.DB) *Repository {
	return &Repository{
		db: pr,
	}
}

//Repository - db
type Repository struct {
	db *sqlx.DB
}

//PingDB ...
func (r *Repository) PingDB() error {
	err := r.db.Ping()
	if err != nil {
		log.Panicln(err)
	}
	return nil
}
