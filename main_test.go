package main

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type MockProg struct {
	exitCode int
}

func (p *MockProg) Exit(code int) {
	p.exitCode = code
}

func TestRun(t *testing.T) {
	testCases := []struct {
		name     string
		exitCode int
		args     []string
	}{
		{name: "no args", args: []string{}, exitCode: 1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := &MockProg{}
			Run(p)
			assert.Equal(t, p.exitCode, tc.exitCode)
		})
	}
}
