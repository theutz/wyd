package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/theutz/wyd/internal/app"
	"github.com/theutz/wyd/internal/config"
)

func CaptureOutput(t *testing.T, f func()) string {
	t.Helper()

	r, w, err := os.Pipe()
	if err != nil {
		err = fmt.Errorf("error: setting up pipe: %w", err)
		t.Fatal(err)
	}
	defer r.Close()

	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = w
	os.Stderr = w

	f()
	w.Close()

	s := strings.Builder{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s.WriteString(fmt.Sprintf("%s\n", scanner.Text()))
	}
	if scanner.Err() != nil {
		err = fmt.Errorf("error: while scanning output: %w", err)
		t.Fatal(err)
	}
	out := s.String()

	if err != nil {
		t.Fatal(err)
	}

	return out
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
		t.Run(strings.Join(tc.args, " "), func(t *testing.T) {
			// Arrange
			config, err := config.DefaultConfig()
			assert.NoError(t, err)
			config.DatabasePath = ":memory:"
			mockParams := app.NewAppParams{
				Args:           tc.args,
				IsFatalOnError: new(bool),
				MigrationsFS:   &embeddedMigrations,
				Config:         config,
			}
			app := app.NewApp(mockParams)

			// Act
			out := CaptureOutput(t, func() {
				err = app.Run()
				if err != nil {
					logger.Error(err)
					app.Exit(1)
				}
				app.Exit(0)
			})

			// Assert
			cupaloy.SnapshotT(t, out, err)
			assert.Equal(t, tc.exitCode, app.ExitCode())
		})
	}
}
