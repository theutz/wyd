package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Test mg.Namespace

// run all tests
func (Test) All() error {
	err := sh.RunV("gotestsum", "./...")
	if err != nil {
		return mg.Fatal(1, err)
	}
	return nil
}
