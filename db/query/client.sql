-- name: CreateClient :one
INSERT INTO clients (
    owner_login
) VALUES (
    $1
) RETURNING *;