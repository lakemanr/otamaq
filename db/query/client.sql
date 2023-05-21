-- name: CreateClient :one
INSERT INTO clients (
    owner_id
) VALUES (
    $1
) RETURNING *;