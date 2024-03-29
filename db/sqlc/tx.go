package db

import (
	"context"
	"database/sql"
	"fmt"
	"sort"
)

// Sore is the interface for the database.
type Store interface {
	Querier
	CreateOrderTx(ctx context.Context, arg CreateOrderTxParams) (OrderTxResult, error)
}

// SQLStore is a wrapper around sql.DB that implements Store.
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store.
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := store.WithTx(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// DishQty is a dish id and quantity.
type DishQty struct {
	DishID   int32
	Quantity int32
}

// CreateOrderTxParams are the parameters for CreateOrderTx.
type CreateOrderTxParams struct {
	ClientID   int32
	RestID     int32
	DishIDsQty []DishQty
}

// OrderTxResult is the result of CreateOrderTx.
type OrderTxResult struct {
	Order      Order
	OrderItems []OrderItem
}

// CreateOrderTx creates an order and order items in a single database transaction.
func (store *SQLStore) CreateOrderTx(ctx context.Context, arg CreateOrderTxParams) (OrderTxResult, error) {
	var result OrderTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// Create an order.
		result.Order, err = q.CreateOrder(ctx, CreateOrderParams{
			ClientID: arg.ClientID,
			RestID:   arg.RestID,
		})

		if err != nil {
			return err
		}

		// To avoid deadlocks, we sort the dish ids in ascending order.
		sort.Slice(arg.DishIDsQty, func(i, j int) bool {
			return arg.DishIDsQty[i].DishID < arg.DishIDsQty[j].DishID
		})

		// Check unique dish ids.
		for i := 0; i < len(arg.DishIDsQty)-1; i++ {
			if arg.DishIDsQty[i].DishID == arg.DishIDsQty[i+1].DishID {
				return fmt.Errorf("dish id %d is not unique", arg.DishIDsQty[i].DishID)
			}
		}

		// Create order items.
		var orderItem OrderItem
		for _, dishQty := range arg.DishIDsQty {
			orderItem, err = q.CreateOrderItem(ctx, CreateOrderItemParams{
				OrderID:  result.Order.ID,
				DishID:   dishQty.DishID,
				Quantity: dishQty.Quantity,
			})

			if err != nil {
				return err
			}

			result.OrderItems = append(result.OrderItems, orderItem)

			// Decrease the amount of the dish.
			var _ Dish
			_, err = q.AddDishAmount(ctx, AddDishAmountParams{
				ID:     dishQty.DishID,
				Amount: -dishQty.Quantity,
			})

			if err != nil {
				return err
			}
		}
		return nil
	})

	return result, err

}
