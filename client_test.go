package main

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/bradleyjkemp/cupaloy"
)

func Test_AddClient(t *testing.T) { //nolint:paralleltest
	// Arrange
	defer cleanup(t)

	run := func(args ...string) (string, int, error) {
		return runMockApp(t, embeddedMigrations, args...)
	}

	// Act
	out, exitCode, err := run("client", "add", "-n", "Delegator")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 0, exitCode)

	listOut, _, err := run("client", "list")
	assert.NoError(t, err)
	cupaloy.SnapshotT(t, out, listOut)
}

func Test_ListClient(t *testing.T) { //nolint:paralleltest
	// Arrange
	defer cleanup(t)

	run := func(args ...string) (string, int, error) {
		return runMockApp(t, embeddedMigrations, args...)
	}

	names := []string{"Delegator", "Something"}
	for _, name := range names {
		_, _, err := run("client", "add", "-n", name)
		assert.NoError(t, err)
	}

	// Act
	out, exitCode, err := run("client", "list")
	assert.NoError(t, err)
	assert.Equal(t, 0, exitCode)
	cupaloy.SnapshotT(t, out)
}

func Test_DeleteClient(t *testing.T) { //nolint:paralleltest
	// Arrange
	defer cleanup(t)

	run := func(args ...string) (string, int, error) {
		return runMockApp(t, embeddedMigrations, args...)
	}

	names := []string{"Delegator", "Something"}
	for _, name := range names {
		_, _, err := run("client", "add", "-n", name)
		assert.NoError(t, err)
	}

	initialOut, _, err := run("client", "list")
	assert.NoError(t, err)

	// Act
	out, exitCode, err := run("client", "remove", "-n", "Delegator")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 0, exitCode)

	listOut, _, err := run("client", "list")
	assert.NoError(t, err)

	cupaloy.SnapshotT(t, initialOut, out, listOut)
}
