package provider

import (
	"fmt"
	"log"

	"github.com/nikitalier/tenderMonitoring/config"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

// const (
// 	host     = "0.0.0.0"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "mysecretpassword"
// 	dbname   = "postgres"
// 	driver   = "postgres"
// )

// Provider interface for database
type Provider interface {
	Open() error
	GetCon() *sqlx.DB
}

type provider struct {
	DB     *sqlx.DB
	driver string
	info   string
}

//New - new db connection
func New(db *config.SQLDataBase) Provider {
	var info string

	info = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable", db.Server, db.Port, db.User, db.Password, db.DataBase, db.SearchPath)

	return &provider{
		DB:     nil,
		driver: db.Driver,
		info:   info,
	}
}

func (p *provider) GetCon() *sqlx.DB {
	return p.DB
}

//Open is steam id = 94895003
func (p *provider) Open() (err error) {
	p.DB, err = sqlx.Open(p.driver, p.info)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
