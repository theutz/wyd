package cli

import (
	"fmt"
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

func Test_Flag_Help(t *testing.T) {
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

func Test_Flag_Debug(t *testing.T) {
	// Arrange
	p := &MockProgram{}
	c := New(p)

	// Act
	err := c.Run("--debug")

	// Assert
	assert.Error(t, err)
	assert.True(t, c.Value().Debug)
}

func Test_Flag_DatabasePath(t *testing.T) {
	// Arrange
	p := &MockProgram{}
	c := New(p)
	path := "~/.local/share/wyd/custom.db"
	flag := fmt.Sprintf("--database-path=%s", path)

	// Act
	err := c.Run(flag)

	// Assert
	assert.Error(t, err)
	assert.EqualError(t, err, "running kong: no command selected")
	assert.Equal(t, c.Value().DatabasePath, path)
}
