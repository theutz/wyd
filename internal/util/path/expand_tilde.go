package path

import (
	"fmt"
	"os"
	"path/filepath"
)

func ExpandTilde(path string) (string, error) {
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
