package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

//New repository - add db in repository
func New(pr *sqlx.DB, logger zerolog.Logger) *Repository {
	return &Repository{
		db:     pr,
		logger: logger,
	}
}

//Repository - db
type Repository struct {
	db     *sqlx.DB
	logger zerolog.Logger
}

//PingDB - check db connection
func (r *Repository) PingDB() error {
	err := r.db.Ping()
	if err != nil {
		r.logger.Panic().Msg(err.Error())
	}
	return nil
}
