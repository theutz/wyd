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
	args := os.Args
	exited := false
	mockExiter := func(code int) {
		exited = true
		panic(true)
	}
	defer func() { os.Args = args }()

	w := bytes.NewBuffer(nil)
	os.Args = []string{"--help"}
	panicsTrue(t, func() {
		err := run(w, w, mockExiter)
		assert.NoError(t, err)
	})
	assert.True(t, exited)
	assert.Contains(t, w.String(), "Whatch'ya doin'?")
}
