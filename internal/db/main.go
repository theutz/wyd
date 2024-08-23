package db

import (
	"database/sql"

	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func New(file string, l *log.Logger) *sql.DB {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		l.Fatal(err)
	}
	return db
}
