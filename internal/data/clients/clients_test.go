package clients

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

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
