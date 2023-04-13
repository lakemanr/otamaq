package db

import (
	"context"
	"math"
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
	dish, err := testStore.CreateDish(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, dish)

	assert.NotZero(t, dish.ID)
	assert.NotZero(t, dish.CreatedAt)

	assert.Equal(t, arg.RestID, dish.RestID)
	assert.Equal(t, arg.Name, dish.Name)

	return dish
}

func addRandomDishAmount(t *testing.T, dishBefore Dish, amount int32) Dish {

	qtyBefore := dishBefore.Quantity

	arg := AddDishAmountParams{
		Amount: amount,
		ID:     dishBefore.ID,
	}

	dish, err := testStore.AddDishAmount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, dish)

	assert.Equal(t, dishBefore.ID, dish.ID)
	assert.Equal(t, dishBefore.Name, dish.Name)
	assert.Equal(t, dishBefore.RestID, dish.RestID)
	assert.Equal(t, dishBefore.CreatedAt, dish.CreatedAt)

	assert.Equal(t, dish.Quantity, qtyBefore+amount)

	return dish
}

func makeUnlimitedDishAmount(t *testing.T, dishBefore Dish) Dish {
	unlimited := math.MaxInt32 - dishBefore.Quantity
	return addRandomDishAmount(t, dishBefore, unlimited)
}

func TestCreateDish(t *testing.T) {
	rest := createRandomRestaurant(t)
	createRandomDish(t, rest)
}

func TestAddDishAmmount(t *testing.T) {
	rest := createRandomRestaurant(t)
	dishBefore := createRandomDish(t, rest)
	dish := addRandomDishAmount(t, dishBefore, util.RandomQuantity())

	invalidAmount := -(dish.Quantity + util.RandomQuantity())

	arg := AddDishAmountParams{
		Amount: invalidAmount,
		ID:     dishBefore.ID,
	}

	dish, err := testStore.AddDishAmount(context.Background(), arg)
	require.Error(t, err)
	require.Empty(t, dish)
}

func TestGetDish(t *testing.T) {
	rest := createRandomRestaurant(t)
	dishBefore := createRandomDish(t, rest)

	dish, err := testStore.GetDish(context.Background(), dishBefore.ID)
	require.NoError(t, err)
	require.NotEmpty(t, dish)

	assert.Equal(t, dishBefore.ID, dish.ID)
	assert.Equal(t, dishBefore.Name, dish.Name)
	assert.Equal(t, dishBefore.RestID, dish.RestID)
	assert.Equal(t, dishBefore.CreatedAt, dish.CreatedAt)
	assert.Equal(t, dishBefore.Quantity, dish.Quantity)
}
