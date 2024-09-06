package db

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestMakeDsn(t *testing.T) {
	suffix := "foreign_keys=on&journal_mode=WAL"
	testCases := []struct {
		name  string
		path  func() string
		wants string
	}{
		{
			name: "absolute path",
			path: func() string {
				base := "/home/dude/wheres/my/car.db"
				tmpPath := filepath.Join(os.TempDir(), base)
				return tmpPath
			},
			wants: fmt.Sprintf(
				"file:%s?%s",
				filepath.Join(
					os.TempDir(),
					"/home/dude/wheres/my/car.db",
				),
				suffix,
			),
		},
		{
			name: "in-memory database",
			path: func() string {
				return ":memory:"
			},
			wants: fmt.Sprintf(":memory:?%s", suffix),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			got, err := makeDsn(tc.path())

			// Assert
			assert.NoError(t, err)
			assert.Equal(t, tc.wants, got)
		})
	}
}

func TestNew(t *testing.T) {
	// Arrange
	ctx := context.Background()
	path := ":memory:"

	// Act
	db, err := New(ctx, path)
	assert.NoError(t, err)
	defer db.Close()

	// Assert
	assert.NotZero(t, db)
}

func TestBasicQuery(t *testing.T) {
	// Arrange
	ctx := context.Background()
	path := ":memory:"
	db, err := New(ctx, path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	q := db.Queries()

	// Act
	project, err := q.AddProject(ctx, "boo")
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, project.Name, "boo")
}
