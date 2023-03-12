-- name: CreateDish :one
INSERT INTO dishes (
    name, 
    rest_id
) VALUES (
    $1, $2
) RETURNING *;
