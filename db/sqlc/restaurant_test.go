package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateRestaurant(t *testing.T) {
	arg := "Отличный Ресторан"
	rest, err := testQueries.CreateRestaurant(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, rest)
	assert.Equal(t, rest.Name, arg)
	assert.NotZero(t, rest.ID)
	assert.NotZero(t, rest.CreatedAt)
}
