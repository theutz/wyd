package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func panicsTrue(t *testing.T, f func()) {
	defer func() {
		if value := recover(); value != nil {
			if boolval, ok := value.(bool); !ok || !boolval {
				t.Fatalf("expected panic with true but got %v", value)
			}
		}
	}()
	f()
	t.Fatal("expected panic did not occur")
}

func TestRun(t *testing.T) {
	tests := []struct {
		name  string
		args  []string
		wants string
	}{
		{"Help flag", []string{"--help"}, "Whatch'ya doin'?"},
		{"Version flag", []string{"--version"}, "unknown (built from source)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldArgs := os.Args
			exited := false
			mockExiter := func(code int) {
				exited = true
				panic(true)
			}
			defer func() { os.Args = oldArgs }()

			w := bytes.NewBuffer(nil)
			args := append([]string{"wyd"}, tt.args...)
			os.Args = args
			panicsTrue(t, func() {
				err := run(w, w, mockExiter)
				assert.NoError(t, err)
			})
			assert.True(t, exited)
			assert.Contains(t, w.String(), tt.wants)
		})
	}
}
