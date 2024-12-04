-- name: GetReceipt :one
SELECT * FROM receipts
WHERE id = $1;

-- name: CreateReceipt :one
INSERT INTO receipts (
    retailer, purchase_date, purchase_time, total, points
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetReceiptByUuid :one
SELECT * FROM receipts 
WHERE uuid = $1;