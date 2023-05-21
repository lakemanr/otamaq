-- name: CreateRestaurant :one
INSERT INTO restaurants (
    owner_id,
    name
) VALUES (
    $1, $2
) RETURNING *;

-- name: ListRestaurants :many
SELECT * FROM restaurants
ORDER BY id
LIMIT $1
OFFSET $2;