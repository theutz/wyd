package main

import (
	"context"
	"embed"
	"os"
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/theutz/wyd/internal/app"
	"github.com/theutz/wyd/internal/config"
	"github.com/theutz/wyd/internal/util"
)

const testDbPath = "db.sqlite"

//////////////////
// Root Command //
//////////////////

func Test_NoArgs(t *testing.T) { //nolint:paralleltest
	out, exitCode, err := runMockApp(t, embeddedMigrations, "")
	defer cleanup(t)

	// Assert
	cupaloy.SnapshotT(t, out, err, exitCode)
}

func Test_Help(t *testing.T) { //nolint:paralleltest
	testCases := []struct {
		args []string
	}{
		{[]string{}},
		{[]string{"--help"}},
		{[]string{"config", "--help"}},
		{[]string{"client", "--help"}},
		{[]string{"client", "list", "--help"}},
		{[]string{"client", "add", "--help"}},
	}

	for _, testCase := range testCases { //nolint:paralleltest
		t.Run(strings.Join(testCase.args, " "), func(t *testing.T) {
			out, exitCode, err := runMockApp(t, embeddedMigrations, testCase.args...)
			defer cleanup(t)

			// Assert
			cupaloy.SnapshotT(t, out, err, exitCode)
		})
	}
}

////////////////////
// Config Command //
////////////////////

func Test_ConfigShow(t *testing.T) { //nolint:paralleltest
	// Arrange
	defer cleanup(t)

	run := func(args ...string) (string, int, error) {
		return runMockApp(t, embeddedMigrations, args...)
	}

	// Act
	out, exitCode, err := run("config", "show")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 0, exitCode)

	assert.NoError(t, err)
	cupaloy.SnapshotT(t, out)
}

////////////////////
// Client Command //
////////////////////

func Test_AddClient(t *testing.T) { //nolint:paralleltest
	// Arrange
	defer cleanup(t)

	run := func(args ...string) (string, int, error) {
		return runMockApp(t, embeddedMigrations, args...)
	}

	// Act
	out, exitCode, err := run("client", "add", "-n", "New Client")

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

	names := []string{"New Client", "Other New Client"}
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

	names := []string{"New Client", "Other New Client"}
	for _, name := range names {
		_, _, err := run("client", "add", "-n", name)
		assert.NoError(t, err)
	}

	initialOut, _, err := run("client", "list")
	assert.NoError(t, err)

	// Act
	out, exitCode, err := run("client", "remove", "-n", "New Client")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 0, exitCode)

	listOut, _, err := run("client", "list")
	assert.NoError(t, err)

	cupaloy.SnapshotT(t, initialOut, out, listOut)
}

/////////////////////
// Project Command //
/////////////////////

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

/////////////
// Helpers //
/////////////

func runMockApp(t *testing.T, migrationFS embed.FS, args ...string) (string, int, error) {
	t.Helper()

	config, err := config.DefaultConfig()
	assert.NoError(t, err)

	config.DatabasePath = testDbPath
	ctx := context.Background()

	isFatalOnError := false
	mockParams := app.NewAppParams{
		Args:           args,
		IsFatalOnError: &isFatalOnError,
		MigrationsFS:   &migrationFS,
		Config:         config,
		Context:        &ctx,
	}

	app := app.NewApp(mockParams)
	out, err := util.CaptureOutput(func() error {
		err := app.Run()
		if err != nil {
			app.Exit(1)
		}

		app.Exit(0)

		return err //nolint:wrapcheck
	})

	return out, app.ExitCode(), err
}

func cleanup(t *testing.T) {
	t.Helper()

	err := os.Remove(testDbPath)
	if err != nil {
		t.Fatal(err)
	}
}

// vim: fdl=0
