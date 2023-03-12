package db

import (
	"context"
	"testing"

	"github.com/lakemanr/otamaq/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomRestaurant(t *testing.T) Restaurant {
	arg := util.RandomRestaurantName()
	rest, err := testQueries.CreateRestaurant(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, rest)
	assert.Equal(t, rest.Name, arg)
	assert.NotZero(t, rest.ID)
	assert.NotZero(t, rest.CreatedAt)

	return rest
}

func TestCreateRestaurant(t *testing.T) {
	createRandomRestaurant(t)
}
