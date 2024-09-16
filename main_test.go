package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

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
		{"no args", []string{}, 1},
		{"help flag", []string{"--help"}, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			mockLogger := log.New(io.Discard)
			mockApp := app.NewAppParams{
				Logger: mockLogger,
				Args:   tc.args,
			}

			// Act
			out := CaptureOutput(t, func() {
				app.NewApp(mockApp)
			})

			// Assert
			cupaloy.SnapshotT(t, out)
		})
	}
}
