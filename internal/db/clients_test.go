package db

import (
	"context"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/jaswdr/faker"
	"github.com/theutz/wyd/internal/db/queries"
)

func mkClients(t *testing.T, count int) (context.Context, *queries.Queries, []queries.Client) {
	t.Helper()
	ctx, q := mkQ(t)

	clients := []queries.Client{}
	faker := faker.New()

	for i := 0; i < count; i++ {
		name := faker.Person().Name()

		c, err := q.AddClient(ctx, name)
		if err != nil {
			t.Fatal(err)
		}

		clients = append(clients, c)
	}

	return ctx, q, clients
}

func TestAddClient(t *testing.T) {
	// Arrange
	ctx, q := mkQ(t)

	// Act
	client, err := q.AddClient(ctx, "delegator")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "delegator", client.Name)
}

func TestListClients(t *testing.T) {
	// Arrange
	ctx, q, _ := mkClients(t, 5)

	// Act
	clients, err := q.ListClients(ctx)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 5, len(clients))
}

func TestDeleteClient(t *testing.T) {
	// Arrange
	count := 20
	ctx, q, clients := mkClients(t, count)

	// Act
	_, err := q.DeleteClient(ctx, clients[1].Name)

	// Assert
	assert.NoError(t, err)
	clients, err = q.ListClients(ctx)
	assert.NoError(t, err)
	assert.Equal(t, count-1, len(clients))
}
