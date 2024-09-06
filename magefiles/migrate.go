package main

import (
	"github.com/huandu/xstrings"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Migrate mg.Namespace

var migrationsDir = "internal/db/migrations"

// create a new migration with `name`
func (Migrate) Create(name string) error {
	name = xstrings.ToSnakeCase(name)
	err := sh.Run(
		"goose",
		"-dir",
		migrationsDir,
		"create",
		name,
		"sql",
	)
	if err != nil {
		return err
	}

	return nil
}
