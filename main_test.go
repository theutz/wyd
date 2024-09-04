package main

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name     string
		exitCode int
		args     []string
	}{
		{name: "no args", args: []string{}, exitCode: 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var got int
			run(func(code int) {
				got = code
			})
			assert.Equal(t, got, tc.exitCode)
		})
	}
}
