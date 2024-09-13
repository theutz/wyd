package cli

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/bradleyjkemp/cupaloy"
	"github.com/theutz/wyd/internal/utils"
)

type MockProgram struct {
	exitCode int
}

func (p *MockProgram) Exit(code int) {
	p.exitCode = code
}

type MockCli struct{}

func TestHelpFlag(t *testing.T) {
	testCases := []struct {
		name string
		args []string
	}{
		{"long", []string{"--help"}},
		{"short", []string{"-h"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			p := &MockProgram{
				exitCode: -1,
			}
			c := New(p)

			// Act
			out, err := utils.CaptureOutput(t, func() error {
				return c.Run(tc.args...)
			})

			// Assert
			assert.Error(t, err)
			assert.Equal(t, p.exitCode, 0)
			cupaloy.SnapshotT(t, out, err)
		})
	}
}

func TestDebugFlag(t *testing.T) {
	testCases := []struct {
		name  string
		args  []string
		wants bool
	}{
		{
			name:  "long",
			args:  []string{"--verbose"},
			wants: true,
		},
		{
			name:  "short",
			args:  []string{"-v"},
			wants: true,
		},
		{
			name:  "none",
			args:  []string{},
			wants: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			p := &MockProgram{}
			c := New(p)

			// Act
			err := c.Run(tc.args...)

			// Assert
			assert.Error(t, err)
			assert.Equal(t, tc.wants, c.Cmd().Debug)
			cupaloy.SnapshotT(t, err)
		})
	}
}

func TestDatabasePathFlag(t *testing.T) {
	flag := "--database-path"
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		panic("couldn't get current file path")
	}
	configPath := filepath.Clean(filepath.Join(filepath.Dir(currentFile), "app", "config.yml"))

	testCases := []struct {
		name  string
		args  []string
		wants string
	}{
		{
			name:  "default value",
			args:  []string{fmt.Sprintf("%s=%s", flag, configPath)},
			wants: configPath,
		},
		{
			name:  "absolute path",
			args:  []string{fmt.Sprintf("%s=%s", flag, currentFile)},
			wants: currentFile,
		},
		{
			name:  "no path",
			args:  []string{fmt.Sprintf("%s=", flag)},
			wants: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			p := &MockProgram{}
			c := New(p)

			// Act
			err := c.Run(tc.args...)

			// Assert
			assert.Equal(t, tc.wants, c.Cmd().DatabasePath)
			cupaloy.SnapshotT(t, err)
		})
	}
}
