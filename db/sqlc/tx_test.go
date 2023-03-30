package db

import (
	"context"
	"testing"

	"github.com/lakemanr/otamaq/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createRandomOrderTx(t *testing.T, client Client, rest Restaurant, dishes []Dish) (OrderTxResult, error) {

	dishIDsQty := make(map[int32]int32)

	for _, dish := range dishes {
		dishIDsQty[dish.ID] = util.RandomQuantity()
	}

	arg := CreateOrderTxParams{
		ClientID:   client.ID,
		RestID:     rest.ID,
		DishIDsQty: dishIDsQty,
	}

	order, err := CreateOrderTx(testDb, context.Background(), arg)

	return order, err
}

func TestCreateOrderTx(t *testing.T) {

	client := createRandomClient(t)
	rest := createRandomRestaurant(t)

	orderSize := util.RandomOrderSize()
	dishes := make([]Dish, orderSize)
	dishIDs := make([]int32, orderSize)

	for i := 0; i < orderSize; i++ {
		dish := createRandomDish(t, rest)
		dishes[i] = dish
		dishIDs[i] = dish.ID
	}

	errs := make(chan error)
	orders := make(chan OrderTxResult)

	numOrders := util.RandomNumOrders()

	for i := 0; i < numOrders; i++ {
		go func() {
			order, err := createRandomOrderTx(t, client, rest, dishes)
			errs <- err
			orders <- order
		}()
	}

	for i := 0; i < numOrders; i++ {
		err := <-errs
		require.NoError(t, err)

		order := <-orders
		require.NotEmpty(t, order)

		assert.NotZero(t, order.Order.ID)
		assert.NotZero(t, order.Order.CreatedAt)

		assert.Equal(t, client.ID, order.Order.ClientID)
		assert.Equal(t, rest.ID, order.Order.RestID)

		assert.Equal(t, len(dishes), len(order.OrderItems))

		for _, orderItem := range order.OrderItems {
			assert.NotZero(t, orderItem.ID)
			assert.Equal(t, order.Order.ID, orderItem.OrderID)
			assert.NotZero(t, orderItem.Quantity)
			assert.Contains(t, dishIDs, orderItem.DishID)
		}
	}

}
