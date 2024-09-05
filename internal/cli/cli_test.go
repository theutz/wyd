package cli

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type MockProgram struct {
	exitCode int
}

func (p *MockProgram) Exit(code int) {
	p.exitCode = code
}

type MockCli struct{}

func TestHelpFlag(t *testing.T) {
	// Arrange
	p := &MockProgram{
		exitCode: -1,
	}
	c := New(p)

	// Act
	err := c.Run("--help")

	// Assert
	assert.Error(t, err)
	assert.Equal(t, p.exitCode, 0)
}

func TestDebugFlag(t *testing.T) {
	// Arrange
	p := &MockProgram{}
	c := New(p)

	// Act
	err := c.Run("--debug")

	// Assert
	assert.Error(t, err)
	assert.True(t, c.Value().Debug)
}
