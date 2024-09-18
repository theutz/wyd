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

	mockParams := app.NewAppParams{
		Args:           args,
		IsFatalOnError: new(bool),
		MigrationsFS:   &migrationFS,
		Config:         config,
		Context:        &ctx,
	}

	app := app.NewApp(mockParams)
	out, err := util.CaptureOutput(func() error {
		err := app.Run()
		if err != nil {
			return err //nolint:wrapcheck
		}

		return nil
	})

	return out, app.ExitCode(), err
}

func Test_Help(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		args     []string
		exitCode int
	}{
		{[]string{}, 0},
		{[]string{"--help"}, 0},
		{[]string{"config", "--help"}, 0},
		{[]string{"config", "show"}, 0},
		{[]string{"client"}, 0},
		{[]string{"client", "--help"}, 0},
		{[]string{"client", "list"}, 0},
		{[]string{"client", "list", "--help"}, 0},
		{[]string{"client", "add", "--help"}, 0},
	}

	for _, testCase := range testCases {
		os.Remove(testDbPath)
		t.Run(strings.Join(testCase.args, " "), func(t *testing.T) {
			t.Parallel()
			out, exitCode, err := RunMockApp(t, embeddedMigrations, testCase.args...)

			// Assert
			cupaloy.SnapshotT(t, out, err)
			assert.Equal(t, testCase.exitCode, exitCode)
		})
	}
}

func Test_AddClient(t *testing.T) {
	t.Parallel()

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
