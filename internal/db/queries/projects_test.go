package queries

import (
	"context"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/theutz/wyd/internal/db"
)

func TestAddProject(t *testing.T) {
	// Arrange
	ctx := context.Background()
	db, err := db.New(ctx, ":memory:")
	assert.NoError(t, err)
	defer db.Close()
	q := New(db)

	// Act
	project, err := q.AddProject(ctx, AddProjectParams{
		Name:     "boo",
		ClientID: 1,
	})
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, "boo", project.Name)
	assert.Equal(t, 1, project.ClientID)
}
