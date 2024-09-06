package db

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestMakeDsn(t *testing.T) {
	testCases := []struct {
		name   string
		path   string
		suffix string
	}{
		{
			name: "absolute path",
			path: "/home/dude/wheres/my/car.db",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			suffix := "foreign_keys=on&journal_mode=WAL"
			tmpPath := filepath.Join(os.TempDir(), tc.path)
			wants := fmt.Sprintf("file:%s?%s", tmpPath, suffix)

			// Act
			got, err := makeDsn(tmpPath)

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, wants, got)
		})
	}
}

func TestNew(t *testing.T) {
	// Arrange
	file, err := os.CreateTemp(os.TempDir(), "wyd.*****.db")
	if err != nil {
		t.Fatal(err)
	}
	file.Close()
	defer os.Remove(file.Name())
	path := filepath.Join(file.Name())

	// Act
	db, err := New(path)

	// Assert
	assert.NoError(t, err)
	assert.NotZero(t, db)
}
