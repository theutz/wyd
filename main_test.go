package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/utils"
)

type MockProgram program

func (p *MockProgram) Args() []string {
	return p.args
}

func (p *MockProgram) ExitCode() int {
	return p.exitCode
}

func (p *MockProgram) Logger() *log.Logger {
	return p.logger
}

type MockCli struct{}

func (c *MockCli) Run(args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("no args")
	}
	return nil
}

func (c *MockCli) SetConfigPath(path string) {
}

func (c *MockCli) Cmd() RootCmd {
	return RootCmd{}
}

func TestRun(t *testing.T) {
	testCases := []struct {
		name     string
		exitCode int
		args     []string
	}{
		{
			name:     "with no args",
			args:     []string{},
			exitCode: 1,
		},
		{
			name:     "with help flag",
			args:     []string{"--help"},
			exitCode: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := log.New(os.Stderr)
			p := &MockProgram{
				args:     tc.args,
				logger:   logger,
				exitCode: -1,
			}
			c := &MockCli{}

			// Act
			out, err := utils.CaptureOutput(t, func() error {
				Init(p, c)
				return nil
			})

			// Assert
			assert.NoError(t, err)
			cupaloy.SnapshotT(t, out)
			assert.Equal(t, tc.exitCode, p.ExitCode())
		})
	}
}
