-- name: CreateDish :one
INSERT INTO dishes (
    name, 
    rest_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetDish :one
SELECT * FROM dishes WHERE id = $1;

-- name: AddDishAmount :one
UPDATE dishes 
SET quantity = quantity + sqlc.arg(amount) 
WHERE id = sqlc.arg(id) 
RETURNING *;
