// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: items.sql

package db

import (
	"context"
)

const createItem = `-- name: CreateItem :one
INSERT INTO items (
    receipt_id, short_description, price
) VALUES (
    $1, $2, $3
)
RETURNING id, receipt_id, short_description, price
`

type CreateItemParams struct {
	ReceiptID        int32
	ShortDescription string
	Price            string
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, createItem, arg.ReceiptID, arg.ShortDescription, arg.Price)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.ReceiptID,
		&i.ShortDescription,
		&i.Price,
	)
	return i, err
}

const getItem = `-- name: GetItem :one
SELECT id, receipt_id, short_description, price FROM items
WHERE id = $1
`

func (q *Queries) GetItem(ctx context.Context, id int32) (Item, error) {
	row := q.db.QueryRowContext(ctx, getItem, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.ReceiptID,
		&i.ShortDescription,
		&i.Price,
	)
	return i, err
}

const getReceiptItems = `-- name: GetReceiptItems :many
SELECT id, receipt_id, short_description, price FROM items
WHERE receipt_id = $1
`

func (q *Queries) GetReceiptItems(ctx context.Context, receiptID int32) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, getReceiptItems, receiptID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.ReceiptID,
			&i.ShortDescription,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
