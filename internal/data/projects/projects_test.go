package projects

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

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
