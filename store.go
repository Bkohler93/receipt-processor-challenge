package main

import (
	"context"
	"database/sql"

	"github.com/bkohler93/receipt-processor-challenge/db"
	"github.com/google/uuid"
)

type store struct {
	*db.Queries
}

func newStore(database *sql.DB) store {
	q := db.New(database)
	return store{q}
}

func (s *store) AddReceipt(r receipt) (db.Receipt, error) {
	return s.CreateReceipt(context.Background(), db.CreateReceiptParams{
		Retailer:     r.Retailer,
		PurchaseDate: r.PurchaseDate,
		PurchaseTime: r.PurchaseTime,
		Total:        r.Total,
		Points:       int32(r.Points),
	})
}

func (s *store) GetReceipt(uuid uuid.UUID) (receipt, error) {
	dr, err := s.GetReceiptByUuid(context.Background(), uuid)

	// if extending this application items should be retrieved from the database that have matching receipt id

	if err != nil {
		return receipt{}, err
	}

	return fromDatabaseReceipt(dr, nil), nil //pass 'nil' in for items slice because it is unneeded for the challenge
}
