-- name: CreateClient :one
INSERT INTO clients (
    full_name,
    login
) VALUES (
    $1, $2
) RETURNING *;