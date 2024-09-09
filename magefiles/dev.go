package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Dev mg.Namespace

func (Dev) Up() error {
	args := []string{
		"--wrap-process=none",
		"--restart",
		"--watch=process-compose.yml",
		"--watch=magefiles/dev.go",
		"--",
		"process-compose",
		"--keep-project",
	}
	sh.Run("watchexec", args...)
	return nil
}

func (Dev) Attach() error {
	sh.Run("process-compose", "attach")
	return nil
}
