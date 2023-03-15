package db

import (
	"context"
	"testing"

	"github.com/lakemanr/otamaq/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomClient(t *testing.T) Client {
	arg := CreateClientParams{
		FullName: util.RandomClientName(),
		Login:    util.RandomClientLogin(),
	}

	client, err := testQueries.CreateClient(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, client)

	assert.Equal(t, client.FullName, arg.FullName)
	assert.Equal(t, client.Login, arg.Login)

	assert.NotZero(t, client.ID)
	assert.NotZero(t, client.CreatedAt)

	return client
}

func TestCreateClient(t *testing.T) {
	createRandomClient(t)
}
