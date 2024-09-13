package db

import (
	"context"
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/jaswdr/faker"
	"github.com/theutz/wyd/internal/db/queries"
)

func mkQ(t *testing.T) (ctx context.Context, q *queries.Queries) {
	t.Helper()
	ctx = context.Background()
	db, err := New(ctx, ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	q = queries.New(db)
	return
}

func mkProjects(t *testing.T, count int) (*queries.Queries, context.Context, []queries.Project) {
	t.Helper()
	ctx, q := mkQ(t)
	faker := faker.New()

	projects := []queries.Project{}

	for i := 0; i < count; i++ {
		n := strings.Join(faker.Lorem().Words(faker.IntBetween(1, 4)), " ")

		a := queries.AddProjectParams{Name: n}
		p, err := q.AddProject(ctx, a)
		if err != nil {
			t.Fatal(err)
		}
		projects = append(projects, p)
	}

	return q, ctx, projects
}

func TestAddProject(t *testing.T) {
	// Arrange
	ctx, q := mkQ(t)

	// Act
	project, err := q.AddProject(ctx, queries.AddProjectParams{
		Name:     "boo",
		ClientID: 1,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "boo", project.Name)
	assert.Equal(t, 1, project.ClientID)
}

func TestListProjects(t *testing.T) {
	// Arrange
	q, ctx, wants := mkProjects(t, 4)

	// Act
	projects, err := q.ListProjects(ctx)

	// Assert
	assert.NoError(t, err)
	assert.NotZero(t, projects)
	assert.Equal(t, 4, len(projects))
	for i, p := range projects {
		assert.Equal(t, wants[i], p)
	}
}

func TestDeleteProject(t *testing.T) {
	// Arrange
	q, ctx, projects := mkProjects(t, 2)

	// Act
	p, err := q.DeleteProject(ctx, 1)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 1, p.ID)
	projects, err = q.ListProjects(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(projects))
}
