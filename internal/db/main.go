package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/theutz/wyd/internal/logger"
)

var db *sql.DB

func New(file string, debug bool) *sql.DB {
	log := logger.New(debug)
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
