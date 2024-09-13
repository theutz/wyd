package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func ExpandTildeToHome(path string) (string, error) {
	if path[0:2] == "~/" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("expanding tilde: %w", err)
		}

		path = path[2:]
		path = filepath.Clean(path)
		path = filepath.Join(homeDir, path)
		return path, nil
	}

	return path, nil
}

func CaptureOutput(t *testing.T, f func() error) (string, error) {
	t.Helper()

	r, w, err := os.Pipe()
	if err != nil {
		err = fmt.Errorf("error: setting up pipe: %w", err)
		t.Fatal(err)
	}
	defer r.Close()

	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = w
	os.Stderr = w

	err = f()
	w.Close()

	s := strings.Builder{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s.WriteString(fmt.Sprintf("%s\n", scanner.Text()))
	}
	if scanner.Err() != nil {
		err = fmt.Errorf("error: while scanning output: %w", err)
		t.Fatal(err)
	}
	out := s.String()

	if err != nil {
		return out, err
	}

	return out, nil
}
