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
		"--hide-disabled",
		"--tui-fs",
	}
	sh.Run("watchexec", args...)
	return nil
}
