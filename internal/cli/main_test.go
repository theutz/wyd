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
	fmt.Printf("code %d", code)
	p.exitCode = code
}

type MockCli struct{}

func TestNew(t *testing.T) {
	p := &MockProgram{}
	got := New(p)
	assert.NotZero(t, got)
}

func TestRun(t *testing.T) {
	p := &MockProgram{
		exitCode: 0,
	}
	c := New(p)
	err := c.Run("--help")
	assert.NoError(t, err)
	assert.Equal(t, p.exitCode, 0)
}
