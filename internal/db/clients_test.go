package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/theutz/wyd/internal/db/queries"
)

func makeQ(t *testing.T) (context.Context, *sql.DB, *queries.Queries) {
	t.Helper()

	ctx := context.Background()
	db, err := New(ctx, ":memory:")
	assert.NoError(t, err)
	q := queries.New(db)

	return ctx, db, q
}

func TestAddClient(t *testing.T) {
	// Arrange
	ctx, db, q := makeQ(t)
	defer db.Close()

	// Act
	c, err := q.AddClient(ctx, "delegator")
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, "delegator", c.Name)
}

func TestListClients(t *testing.T) {
	testCases := []struct {
		name         string
		client_names []string
		wants        []queries.Client
	}{
		{
			name:         "no clients",
			client_names: []string{},
			wants:        nil,
		},
		{
			name:         "one client",
			client_names: []string{"Delegator"},
			wants:        []queries.Client{{ID: 1, Name: "Delegator"}},
		},
		{
			name:         "multiple clients",
			client_names: []string{"Huey", "Dewy", "Louis"},
			wants: []queries.Client{
				{ID: 1, Name: "Huey"},
				{ID: 2, Name: "Dewy"},
				{ID: 3, Name: "Louis"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			ctx, db, q := makeQ(t)
			defer db.Close()
			for _, c := range tc.client_names {
				_, err := q.AddClient(ctx, c)
				assert.NoError(t, err)
			}

			// Act
			c, err := q.ListClients(ctx)
			assert.NoError(t, err)

			// Assert
			assert.Equal(t, tc.wants, c)
			assert.Compare(t, tc.wants, c)
		})
	}
}
