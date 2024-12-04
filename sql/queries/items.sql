-- name: GetItem :one
SELECT * FROM items
WHERE id = $1;

-- name: CreateItem :one
INSERT INTO items (
    receipt_id, short_description, price
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: GetReceiptItems :many
SELECT * FROM items
WHERE receipt_id = $1;