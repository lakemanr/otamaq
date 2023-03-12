package db

import (
	"context"
	"testing"

	"github.com/lakemanr/otamaq/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomDish(t *testing.T, rest Restaurant) Dish {
	arg := CreateDishParams{
		Name:   util.RandomDishName(),
		RestID: rest.ID,
	}
	dish, err := testQueries.CreateDish(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, dish)

	assert.NotZero(t, dish.ID)
	assert.NotZero(t, dish.CreatedAt)

	assert.Equal(t, arg.RestID, dish.RestID)
	assert.Equal(t, arg.Name, dish.Name)

	return dish
}

func TestCreateDish(t *testing.T) {
	rest := createRandomRestaurant(t)
	createRandomDish(t, rest)
}
