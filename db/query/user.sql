-- name: CreateUser :one
INSERT INTO users (
    login,
    full_name,
    hashed_passwords
) VALUES (
    $1, $2, $3
) RETURNING *;