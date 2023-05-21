package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomClient(t *testing.T, user User) Client {

	client, err := testStore.CreateClient(context.Background(), user.ID)

	require.NoError(t, err)
	require.NotEmpty(t, client)

	assert.Equal(t, client.OwnerID, user.ID)

	assert.NotZero(t, client.ID)
	assert.NotZero(t, client.CreatedAt)

	return client
}

func TestCreateClient(t *testing.T) {
	user := createRandomUser(t)
	createRandomClient(t, user)
}
