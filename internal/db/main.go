package db

import (
	clog "github.com/charmbracelet/log"
	"github.com/theutz/wyd/project"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(file string, log *clog.Logger) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	db.AutoMigrate(&project.Project{})

	return db
}
