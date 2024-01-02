package provider

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabase(conf *Config) *sqlx.DB {
	db, err := sqlx.Open("sqlite3", conf.Path)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
