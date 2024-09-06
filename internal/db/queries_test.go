package db

import (
	"context"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/theutz/wyd/internal/db/queries"
)

func db(t *testing.T) (
	db DB,
	ctx context.Context,
	q *queries.Queries,
) {
	t.Helper()
	ctx = context.Background()
	db, err := New(ctx, ":memory:")
	assert.NoError(t, err)
	q = db.Queries()
	return
}

func TestProjectsCount(t *testing.T) {
	// Arrange
	db, ctx, q := db(t)
	defer db.Close()

	// Act
	count, err := q.ProjectsCount(ctx)
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, 0, count)

	// Arrange
	_, err = q.AddProject(ctx, queries.AddProjectParams{
		Name:     "boo",
		ClientID: 1,
	})
	assert.NoError(t, err)

	// Act
	count, err = q.ProjectsCount(ctx)
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, 1, count)
}

func TestAddProject(t *testing.T) {
	// Arrange
	db, ctx, q := db(t)
	defer db.Close()

	// Act
	project, err := q.AddProject(ctx, queries.AddProjectParams{
		Name:     "boo",
		ClientID: 1,
	})
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, "boo", project.Name)
	assert.Equal(t, 1, project.ClientID)
}

func TestAddClient(t *testing.T) {
	// Arrange
	db, ctx, q := db(t)
	defer db.Close()

	// Act
	c, err := q.AddClient(ctx, "delegator")
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, "delegator", c.Name)
}
