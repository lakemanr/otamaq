// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: restaurant.sql

package db

import (
	"context"
)

const createRestaurant = `-- name: CreateRestaurant :one
INSERT INTO restaurants (
    name
) VALUES (
    $1
) RETURNING id, name, created_at
`

func (q *Queries) CreateRestaurant(ctx context.Context, name string) (Restaurant, error) {
	row := q.db.QueryRowContext(ctx, createRestaurant, name)
	var i Restaurant
	err := row.Scan(&i.ID, &i.Name, &i.CreatedAt)
	return i, err
}

const listRestaurants = `-- name: ListRestaurants :many
SELECT id, name, created_at FROM restaurants
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListRestaurantsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListRestaurants(ctx context.Context, arg ListRestaurantsParams) ([]Restaurant, error) {
	rows, err := q.db.QueryContext(ctx, listRestaurants, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Restaurant
	for rows.Next() {
		var i Restaurant
		if err := rows.Scan(&i.ID, &i.Name, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
