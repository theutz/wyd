package main

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/bradleyjkemp/cupaloy"
)

func Test_AddProject(t *testing.T) { //nolint:paralleltest
	// Arrange
	defer cleanup(t)

	run := func(args ...string) (string, int, error) {
		return runMockApp(t, embeddedMigrations, args...)
	}
	clientOut, _, err := run("client", "add", "-n", "New Client")
	assert.NoError(t, err)

	// Act
	out, exitCode, err := run("project", "add", "-n", "New Project", "-c", "New Client")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 0, exitCode)

	listOut, _, err := run("project", "list")
	assert.NoError(t, err)
	cupaloy.SnapshotT(t, clientOut, out, listOut)
}

func Test_ListProject(t *testing.T) { //nolint:paralleltest
	// Arrange
	defer cleanup(t)

	run := func(args ...string) (string, int, error) {
		return runMockApp(t, embeddedMigrations, args...)
	}
	clientOut, _, err := run("client", "add", "-n", "New Client")
	assert.NoError(t, err)

	names := []string{"New Project", "New Other Project"}
	for _, name := range names {
		_, _, err := run("project", "add", "-n", name, "-c", "New Client")
		assert.NoError(t, err)
	}

	// Act
	out, exitCode, err := run("project", "list")
	assert.NoError(t, err)
	assert.Equal(t, 0, exitCode)
	cupaloy.SnapshotT(t, clientOut, out)
}

func Test_DeleteProject(t *testing.T) { //nolint:paralleltest
	// Arrange
	defer cleanup(t)

	run := func(args ...string) (string, int, error) {
		return runMockApp(t, embeddedMigrations, args...)
	}
	clientOut, _, err := run("client", "add", "-n", "New Client")
	assert.NoError(t, err)
	projectOut, _, err := run("project", "add", "-n", "New Project", "-c", "New Client")
	assert.NoError(t, err)
	initialListOut, _, err := run("project", "list")
	assert.NoError(t, err)

	// Act
	out, exitCode, err := run("project", "remove", "-n", "New Project")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 0, exitCode)

	listOut, _, err := run("project", "list")
	assert.NoError(t, err)
	cupaloy.SnapshotT(
		t,
		clientOut,
		projectOut,
		initialListOut,
		out,
		listOut,
	)
}
