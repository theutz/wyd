package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Test mg.Namespace

// run all tests
func (Test) All() error {
	mg.Deps(Tidy)

	err := sh.RunV("gotestsum", "./...")
	if err != nil {
		return mg.Fatal(1, err)
	}

	return nil
}

// output from all list commands
func (Test) State() error {
	mg.Deps(Tidy)

	cmds := [][]string{
		{"clients"},
		{"projects"},
	}

	for _, c := range cmds {
		args := []string{"run", "."}
		args = append(args, c...)
		err := sh.RunV("go", args...)
		if err != nil {
			return mg.Fatal(1, err)
		}
	}

	return nil
}

func Generate() error {
	args := []string{"generate"}
	if mg.Verbose() {
		args = append(args, "-v")
	}
	args = append(args, "./...")

	err := sh.Run("go", args...)
	if err != nil {
		return mg.Fatal(1, err)
	}

	return nil
}

func Tidy() error {
	mg.Deps(Generate)

	args := []string{"mod", "tidy"}
	if mg.Verbose() {
		args = append(args, "-v")
	}

	err := sh.Run("go", args...)
	if err != nil {
		return mg.Fatal(1, err)
	}

	return nil
}
