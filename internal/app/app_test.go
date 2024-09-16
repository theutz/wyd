package app

import (
	"io"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/charmbracelet/log"
)

func TestLogger(t *testing.T) {
	// Arrange
	a := NewApp(NewAppParams{
		Logger: log.New(io.Discard),
		Args:   []string{},
	})

	// Act
	l := a.Logger()

	// Assert
	assert.NotZero(t, l)
}

func TestArgs(t *testing.T) {
	// Arrange
	a := NewApp(NewAppParams{
		Args: []string{"--help"},
	})

	// Act
	r := a.Args()

	// Assert
	assert.Equal(t, []string{"--help"}, r)
}
