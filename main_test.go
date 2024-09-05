package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/charmbracelet/log"
)

type MockProg struct {
	exitCode int
	Prog
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
			p := &MockProg{}
			p.SetArgs(tc.args)
			c := &MockCli{}

			r, w, err := os.Pipe()
			assert.NoError(t, err)
			log := log.New(w)

			// Act
			Run(p, c, log)
			w.Close()

			scanner := bufio.NewScanner(r)
			str := strings.Builder{}
			for scanner.Scan() {
				if _, err := str.Write(scanner.Bytes()); err != nil {
					panic(err)
				}
			}
			if err := scanner.Err(); err != nil {
				panic(err)
			}

			// Assert
			assert.Equal(t, p.exitCode, tc.exitCode)
			assert.Contains(t, str.String(), tc.errMessage)
		})
	}
}
