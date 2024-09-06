package clients

import (
	"context"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/theutz/wyd/internal/db"
)

func TestAddClient(t *testing.T) {
	// Arrange
	ctx := context.Background()
	db, err := db.New(ctx, ":memory:")
	assert.NoError(t, err)
	defer db.Close()
	q := New(db)

	// Act
	c, err := q.AddClient(ctx, "delegator")
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, "delegator", c.Name)
}
