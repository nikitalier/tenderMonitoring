package provider

import (
	"fmt"

	"github.com/nikitalier/tenderMonitoring/config"
	"github.com/rs/zerolog"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

// Provider interface for database
type Provider interface {
	Open() error
	GetCon() *sqlx.DB
}

type provider struct {
	DB     *sqlx.DB
	logger zerolog.Logger
	driver string
	info   string
}

//New - new db connection
func New(db *config.SQLDataBase, logger zerolog.Logger) Provider {
	var info string

	info = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable", db.Server, db.Port, db.User, db.Password, db.DataBase, db.SearchPath)

	return &provider{
		DB:     nil,
		driver: db.Driver,
		info:   info,
		logger: logger,
	}
}

//GetCon - get connection
func (p *provider) GetCon() *sqlx.DB {
	return p.DB
}

//Open new con
func (p *provider) Open() (err error) {
	p.DB, err = sqlx.Open(p.driver, p.info)

	if err != nil {
		p.logger.Error().Msg(err.Error())
		return err
	}

	return nil
}
