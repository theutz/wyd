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
	})

	// Act
	l := a.Logger()

	// Assert
	assert.NotZero(t, l)
}
