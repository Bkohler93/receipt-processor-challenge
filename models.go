package main

import (
	"fmt"
	"regexp"
	"time"

	"github.com/bkohler93/receipt-processor-challenge/db"
)

type receipt struct {
	Retailer     string
	PurchaseDate time.Time
	PurchaseTime time.Time
	Total        string
	Items        []item
	Points       int
}

type item struct {
	ShortDescription string
	Price            string
}

type receiptRequest struct {
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Total        string
	Items        []itemRequest
}

type itemRequest struct {
	ShortDescription string
	Price            string
}

func (rr receiptRequest) validateReceipt() (receipt, error) {
	var r receipt
	invalidFields := []string{}
	isInvalid := false

	retailerPattern := `^[\w\s\-&]+$`
	purchaseDateFormat := `2006-01-02`
	purchaseTimeFormat := `15:04`
	totalPattern := `^\d+\.\d{2}$`

	rPrx := regexp.MustCompile(retailerPattern)
	if !rPrx.MatchString(rr.Retailer) {
		invalidFields = append(invalidFields, "retailer")
		isInvalid = true
	}

	purchaseDate, err := time.Parse(purchaseDateFormat, rr.PurchaseDate)
	if err != nil {
		invalidFields = append(invalidFields, "purchaseDate")
		isInvalid = true
	}

	purchaseTime, err := time.Parse(purchaseTimeFormat, rr.PurchaseTime)
	if err != nil {
		invalidFields = append(invalidFields, "purchaseTime")
		isInvalid = true
	}

	trx := regexp.MustCompile(totalPattern)
	if !trx.MatchString(rr.Total) {
		invalidFields = append(invalidFields, "total")
	}

	items := []item{}
	for in, i := range rr.Items {
		if item, err := i.toItem(in + 1); err != nil {
			invalidFields = append(invalidFields, err.Error())
			isInvalid = true
		} else {
			items = append(items, item)
		}
	}

	if isInvalid {
		return r, fmt.Errorf("invalid receipt request. Invalid fields: %v", invalidFields)
	} else {
		r.Items = items
		r.PurchaseDate = purchaseDate
		r.PurchaseTime = purchaseTime
		r.Retailer = rr.Retailer
		r.Total = rr.Total

		return r, nil
	}
}

func (ir itemRequest) toItem(itemNum int) (item, error) {
	var i item
	invalidFields := []string{}
	isInvalid := false

	shortDescriptionPattern := `^[\w\s\-]+$`
	pricePattern := `^\d+\.\d{2}$`

	sDrx := regexp.MustCompile(shortDescriptionPattern)
	if !sDrx.MatchString(ir.ShortDescription) {
		invalidFields = append(invalidFields, "shortDescription")
		isInvalid = true
	}

	prx := regexp.MustCompile(pricePattern)
	if !prx.MatchString(ir.Price) {
		invalidFields = append(invalidFields, "price")
		isInvalid = true
	}

	if isInvalid {
		return i, fmt.Errorf("item %d - %v", itemNum, invalidFields)
	} else {
		i.Price = ir.Price
		i.ShortDescription = ir.ShortDescription
		return i, nil
	}
}

func fromDatabaseReceipt(dr db.Receipt, items []item) receipt {
	return receipt{
		Retailer:     dr.Retailer,
		PurchaseDate: dr.PurchaseDate,
		PurchaseTime: dr.PurchaseTime,
		Total:        dr.Total,
		Items:        items,
		Points:       int(dr.Points),
	}
}

func (r *receipt) calculatePoints() {
	points := 0

	points += alphanumericCountPoints(r.Retailer)
	points += mustNoCentsPoints(r.Total)
	points += mustIsMultipleOfQuarterPoints(r.Total)
	points += numItemsPoints(r.Items)

	for _, i := range r.Items {
		points += mustTrimmedLengthMultOfThreePoints(i.ShortDescription, i.Price)
	}

	points += oddDayPoints(r.PurchaseDate)
	points += timeBetweenTwoAndFourPoints(r.PurchaseTime)

	r.Points = points
}
