package app

import (
	"io"
	"strconv"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/charmbracelet/log"
)

func TestLogger(t *testing.T) {
	// Arrange
	app := NewApp(NewAppParams{
		Logger: log.New(io.Discard),
		Args:   []string{},
	})

	// Act
	logger := app.Logger()

	// Assert
	assert.NotZero(t, logger)
}

func TestArgs(t *testing.T) {
	// Arrange
	app := NewApp(NewAppParams{
		Args: []string{"--help"},
	})

	// Act
	args := app.Args()

	// Assert
	assert.Equal(t, []string{"--help"}, args)
}

type mockApp App

func (a *mockApp) Exit(code int) {
	a.exitCode = code
}

func (a *mockApp) ExitCode() int {
	return a.exitCode
}

func TestExit(t *testing.T) {
	testCases := []int{0, 1, 130}
	for _, code := range testCases {
		t.Run(strconv.Itoa(code), func(t *testing.T) {
			// Arrange
			app := mockApp{
				logger: log.New(io.Discard),
				args:   []string{},
			}

			// Act
			app.Exit(code)
			got := app.ExitCode()

			// Assert
			assert.Equal(t, code, got)
		})
	}
}
