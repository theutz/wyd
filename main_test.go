package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/app"
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

func TestNewApp(t *testing.T) {
	testCases := []struct {
		name     string
		args     []string
		exitCode int
	}{
		{"no args", []string{}, 0},
		{"help flag", []string{"--help"}, 0},
		{"config help", []string{"config", "--help"}, 0},
		{"show config", []string{"config show"}, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			mockLogger := log.New(os.Stderr)
			mockParams := app.NewAppParams{
				Logger:         mockLogger,
				Args:           tc.args,
				IsFatalOnError: false,
			}
			app := app.NewApp(mockParams)

			var err error
			// Act
			out := CaptureOutput(t, func() {
				err = app.Run()
			})

			// Assert
			cupaloy.SnapshotT(t, out, err)
			assert.Equal(t, tc.exitCode, app.ExitCode())
		})
	}
}
