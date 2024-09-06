package cli

import (
	"fmt"
	"path/filepath"
	"runtime"
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
	testCases := []struct {
		name string
		args []string
		err  string
	}{
		{
			name: "long",
			args: []string{"--help"},
			err:  "parsing kong: expected one of",
		},
		{
			name: "short",
			args: []string{"-h"},
			err:  "parsing kong: expected one of",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			p := &MockProgram{
				exitCode: -1,
			}
			c := New(p)

			// Act
			err := c.Run(tc.args...)

			// Assert
			assert.Error(t, err)
			assert.Contains(t, err.Error(), tc.err)
			assert.Equal(t, p.exitCode, 0)
		})
	}
}

func Test_Flag_Debug(t *testing.T) {
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
			assert.NoError(t, err)
			assert.Equal(t, c.Value().Debug, tc.wants)
		})
	}
}

func Test_Flag_DatabasePath(t *testing.T) {
	flag := "--database-path"
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		panic("couldn't get current file path")
	}
	configPath := filepath.Join(filepath.Dir(currentFile), "testdata", "config.yml")

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
			name: "no path",
			args: []string{fmt.Sprintf("%s=", flag)},
			wants: fmt.Sprintf(
				`parsing kong: %s: "%s" exists but is a directory`,
				flag,
				filepath.Dir(currentFile),
			),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			p := &MockProgram{}
			c := New(p)

			// Act
			err := c.Run(tc.args...)
			if err != nil {
				assert.EqualError(t, err, tc.wants)
				return
			}

			// Assert
			assert.Equal(t, c.Value().DatabasePath, tc.wants)
		})
	}
}
