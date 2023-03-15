package db

import (
	"context"
	"testing"

	"github.com/lakemanr/otamaq/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomOrder(t *testing.T, client Client, rest Restaurant) Order {
	arg := CreateOrderParams{
		ClientID: client.ID,
		RestID:   rest.ID,
	}

	order, err := testQueries.CreateOrder(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, order)

	assert.Equal(t, order.ClientID, arg.ClientID)
	assert.Equal(t, order.RestID, arg.RestID)

	assert.NotZero(t, order.ID)
	assert.NotZero(t, order.CreatedAt)

	return order
}

func createRandomOrderItem(t *testing.T, order Order, dish Dish) OrderItem {
	arg := CreateOrderItemParams{
		OrderID:  order.ID,
		DishID:   dish.ID,
		Quantity: util.RandomQuantity(),
	}

	orderItem, err := testQueries.CreateOrderItem(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, orderItem)

	assert.Equal(t, orderItem.DishID, arg.DishID)
	assert.Equal(t, orderItem.OrderID, arg.OrderID)
	assert.Equal(t, orderItem.Quantity, arg.Quantity)

	assert.NotZero(t, order.ID)
	assert.NotZero(t, order.CreatedAt)

	return orderItem
}

func TestCreateOrder(t *testing.T) {
	client := createRandomClient(t)
	rest := createRandomRestaurant(t)
	createRandomOrder(t, client, rest)
}

func TestCreateOrderItem(t *testing.T) {
	rest := createRandomRestaurant(t)
	dish := createRandomDish(t, rest)
	client := createRandomClient(t)
	order := createRandomOrder(t, client, rest)
	createRandomOrderItem(t, order, dish)
}
