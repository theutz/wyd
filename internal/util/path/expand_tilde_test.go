package path

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestExpandTildeToHome(t *testing.T) {
	homeDir, err := os.UserHomeDir()
	assert.NoError(t, err)

	testCases := []struct {
		name  string
		given string
		wants string
	}{
		{
			name:  "without tilde",
			given: "/home/user/.local/share/wyd/wyd.db",
			wants: "/home/user/.local/share/wyd/wyd.db",
		},
		{
			name:  "with tilde",
			given: "~/.local/share/wyd/wyd.db",
			wants: filepath.Join(homeDir, ".local", "share", "wyd", "wyd.db"),
		},
		{
			name:  "dirty path",
			given: "~/.local/../share/wyd.db",
			wants: filepath.Join(homeDir, "share", "wyd.db"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange

			// Act
			got, err := ExpandTilde(tc.given)
			if err != nil {
				assert.Error(t, err, tc.wants)
				t.Fail()
			}

			// Assert
			assert.Equal(t, tc.wants, got)
		})
	}
}
