-- name: CreateRestaurant :one
INSERT INTO restaurants (
    name
) VALUES (
    $1
) RETURNING *;

-- name: ListRestaurants :many
SELECT * FROM restaurants
ORDER BY id
LIMIT $1
OFFSET $2;