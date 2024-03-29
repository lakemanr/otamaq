package db

import (
	"context"
	"testing"

	"github.com/lakemanr/otamaq/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomRestaurant(t *testing.T, user User) Restaurant {
	arg := CreateRestaurantParams{
		OwnerID: user.ID,
		Name:    util.RandomRestaurantName(),
	}

	rest, err := testStore.CreateRestaurant(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, rest)
	assert.Equal(t, rest.Name, arg.Name)
	assert.Equal(t, rest.OwnerID, arg.OwnerID)
	assert.NotZero(t, rest.ID)
	assert.NotZero(t, rest.CreatedAt)

	return rest
}

func TestCreateRestaurant(t *testing.T) {
	user := createRandomUser(t)
	createRandomRestaurant(t, user)
}
