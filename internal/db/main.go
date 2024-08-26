package db

import (
	clog "github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/client"
	"github.com/theutz/wyd/internal/project"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New(file string, log *clog.Logger) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database", "err", err)
	}

	err = db.AutoMigrate(&project.Project{}, &client.Client{})
	if err != nil {
		log.Fatal("failed to migrate database", "err", err)
	}

	return db
}
