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

func RunMockApp(t *testing.T, migrationFS embed.FS, args ...string) (string, int, error) {
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

func Test_Incomplete(t *testing.T) {
	testCases := []struct {
		args []string
	}{
		{[]string{}},
		{[]string{"config", "show"}},
		{[]string{"client"}},
		{[]string{"client", "list"}},
	}

	for _, testCase := range testCases { //nolint:paralleltest
		os.Remove(testDbPath)
		t.Run(strings.Join(testCase.args, " "), func(t *testing.T) {
			out, exitCode, err := RunMockApp(t, embeddedMigrations, testCase.args...)

			// Assert
			cupaloy.SnapshotT(t, out, err, exitCode)
		})
	}
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
		os.Remove(testDbPath)
		t.Run(strings.Join(testCase.args, " "), func(t *testing.T) {
			out, exitCode, err := RunMockApp(t, embeddedMigrations, testCase.args...)

			// Assert
			cupaloy.SnapshotT(t, out, err, exitCode)
		})
	}
}

func Test_AddClient(t *testing.T) { //nolint:paralleltest
	os.Remove(testDbPath)

	run := func(args ...string) (string, int, error) {
		return RunMockApp(t, embeddedMigrations, args...)
	}

	out, exitCode, err := run("client", "add", "-n", "Delegator")
	assert.NoError(t, err)
	assert.Equal(t, "{1 Delegator}\n", out)
	assert.Equal(t, 0, exitCode)

	out, exitCode, err = run("client", "list")
	assert.NoError(t, err)
	assert.Equal(t, "[{1 Delegator}]\n", out)
	assert.Equal(t, 0, exitCode)
}
