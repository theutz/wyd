package wyd

import (
	"fmt"
	"os"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/cli"
	"github.com/theutz/wyd/internal/utils"
)

type MockProg struct {
	exitCode int
	args     []string
	log      *log.Logger
}

// Args implements Program.
func (p *MockProg) Args() []string {
	return p.args
}

// GetLogger implements Program.
func (p *MockProg) GetLogger() *log.Logger {
	return p.log
}

// SetArgs implements Program.
func (p *MockProg) SetArgs(args []string) {
	p.args = args
}

func (p *MockProg) Exit(code int) {
	p.exitCode = code
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

func (c *MockCli) Cmd() cli.RootCmd {
	return cli.RootCmd{}
}

func TestRun(t *testing.T) {
	testCases := []struct {
		name       string
		exitCode   int
		args       []string
		errMessage string
	}{
		{
			name:       "with no args",
			args:       []string{},
			exitCode:   1,
			errMessage: "no args",
		},
		{
			name:       "with help flag",
			args:       []string{"--help"},
			exitCode:   0,
			errMessage: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			l := log.New(os.Stderr)
			p := &MockProg{
				log: l,
			}
			p.SetArgs(tc.args)
			c := &MockCli{}

			// Act
			out, err := utils.CaptureOutput(t, func() error {
				Run(p, c)
				return nil
			})

			// Assert
			cupaloy.SnapshotT(t, out)
			assert.NoError(t, err)
			assert.Equal(t, p.exitCode, tc.exitCode)
		})
	}
}
