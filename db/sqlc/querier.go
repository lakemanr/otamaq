// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"
)

type Querier interface {
	AddDishAmount(ctx context.Context, arg AddDishAmountParams) (Dish, error)
	CreateClient(ctx context.Context, ownerID int32) (Client, error)
	CreateDish(ctx context.Context, arg CreateDishParams) (Dish, error)
	CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error)
	CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (OrderItem, error)
	CreateRestaurant(ctx context.Context, arg CreateRestaurantParams) (Restaurant, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetDish(ctx context.Context, id int32) (Dish, error)
	ListRestaurants(ctx context.Context, arg ListRestaurantsParams) ([]Restaurant, error)
}

var _ Querier = (*Queries)(nil)
