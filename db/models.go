// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID               int32
	ReceiptID        int32
	ShortDescription string
	Price            string
}

type Receipt struct {
	ID           int32
	Retailer     string
	PurchaseDate time.Time
	PurchaseTime time.Time
	Total        string
	Uuid         uuid.UUID
	Points       int32
}
