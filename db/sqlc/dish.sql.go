// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: dish.sql

package db

import (
	"context"
)

const createDish = `-- name: CreateDish :one
INSERT INTO dishes (
    name, 
    rest_id
) VALUES (
    $1, $2
) RETURNING id, name, rest_id, created_at
`

type CreateDishParams struct {
	Name   string `json:"name"`
	RestID int32  `json:"rest_id"`
}

func (q *Queries) CreateDish(ctx context.Context, arg CreateDishParams) (Dish, error) {
	row := q.db.QueryRowContext(ctx, createDish, arg.Name, arg.RestID)
	var i Dish
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.RestID,
		&i.CreatedAt,
	)
	return i, err
}
