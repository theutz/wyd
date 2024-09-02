package main

import (
	"bytes"
	"io"
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

func TestRun_Subcommands(t *testing.T) {
	tests := []struct {
		name  string
		args  []string
		wants string
	}{
		{"list of projects", []string{"projects", "list"}, "Project Name"},
		{"list of clients", []string{"clients", "list"}, "Client Name"},
		{"list of tasks", []string{"tasks", "list"}, "Task Name"},
		{"list of entries", []string{"entries", "list"}, "Task Name"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			os.Args = append([]string{"wyd"}, tt.args...)

			r, w, err := os.Pipe()
			if err != nil {
				t.Fatal(err)
			}
			defer func() { r.Close() }()

			oldStdout := os.Stdout
			os.Stdout = w
			defer func() { os.Stdout = oldStdout }()

			// Act
			if err = run(w, w, func(c int) {}); err != nil {
				t.Fatal(err)
			}
			w.Close()
			var buf bytes.Buffer
			if _, err = io.Copy(&buf, r); err != nil {
				t.Fatal(err)
			}

			// Assert
			got := buf.String()
			assert.Contains(t, got, tt.wants)
		})
	}
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
