package main

import (
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

var testDbPath = "db.sqlite"

func RunMockApp(t *testing.T, migrationFS embed.FS, args ...string) (string, error, int) {
	config, err := config.DefaultConfig()
	assert.NoError(t, err)

	config.DatabasePath = testDbPath

	mockParams := app.NewAppParams{
		Args:           args,
		IsFatalOnError: new(bool),
		MigrationsFS:   &migrationFS,
		Config:         config,
	}

	app := app.NewApp(mockParams)
	out := util.CaptureOutput(t, func() {
		err = app.Run()
		if err != nil {
			logger.Error(err)
			app.Exit(1)
		}
		app.Exit(0)
	})

	return out, err, app.ExitCode()
}

func Test_Run(t *testing.T) {
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
		{[]string{"client", "add", "--help"}, 0},
		{[]string{"client", "add", "-n", "Delegator"}, 0},
	}

	for _, tc := range testCases {
		os.Remove(testDbPath)
		t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
			out, err, exitCode := RunMockApp(t, embeddedMigrations, tc.args...)

			// Assert
			cupaloy.SnapshotT(t, out, err)
			assert.Equal(t, tc.exitCode, exitCode)
		})
	}
}
