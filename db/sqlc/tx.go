package db

import (
	"context"
	"database/sql"
	"fmt"
)

// execTx executes a function within a database transaction.
func execTx(db *sql.DB, ctx context.Context, fn func(*Queries) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb error: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// CreateOrderTxParams are the parameters for CreateOrderTx.
type CreateOrderTxParams struct {
	ClientID   int32
	RestID     int32
	DishIDsQty map[int32]int32
}

// OrderTxResult is the result of CreateOrderTx.
type OrderTxResult struct {
	Order      Order
	OrderItems []OrderItem
}

// CreateOrderTx creates an order and order items in a single database transaction.
func CreateOrderTx(db *sql.DB, ctx context.Context, arg CreateOrderTxParams) (OrderTxResult, error) {
	var result OrderTxResult

	err := execTx(db, ctx, func(q *Queries) error {
		var err error

		// Create an order.
		result.Order, err = q.CreateOrder(ctx, CreateOrderParams{
			ClientID: arg.ClientID,
			RestID:   arg.RestID,
		})

		if err != nil {
			return err
		}

		// Create order items.
		var orderItem OrderItem
		for dishID, quantity := range arg.DishIDsQty {
			orderItem, err = q.CreateOrderItem(ctx, CreateOrderItemParams{
				OrderID:  result.Order.ID,
				DishID:   dishID,
				Quantity: quantity,
			})

			if err != nil {
				return err
			}

			result.OrderItems = append(result.OrderItems, orderItem)
		}
		return nil
	})

	return result, err

}
