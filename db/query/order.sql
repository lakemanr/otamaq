-- name: createOrder :one
INSERT INTO orders (
    client_id,
    rest_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: createOrderItem :one
INSERT INTO order_items (
    order_id,
    dish_id,
    quantity
) VALUES (
    $1, $2, $3
) RETURNING *;