package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func alphanumericCountPoints(s string) int {
	points := 0
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			points++
		}
	}
	return points
}

// total must be formatted like `99.25` representing a dollar amount
// with variable number of values before the decimal and two values after
func mustNoCentsPoints(total string) int {
	if !prx.MatchString(total) { // prx compiled in `models.go`
		panic(fmt.Errorf("mustNoCentsPoints requires a value formatted as money"))
	}

	cents := strings.Split(total, ".")[1]
	if cents == "00" {
		return 50
	}
	return 0
}

// total must be formatted like `99.25` representing a dollar amount
// with variable number of values before the decimal and two values after
func mustIsMultipleOfQuarterPoints(total string) int {
	if !prx.MatchString(total) { // prx compiled in `models.go`
		panic(fmt.Errorf("mustIsMultipleOfQuarterPoints requires a value formatted as money"))
	}

	cs := strings.Split(total, ".")[1]
	cents, _ := strconv.Atoi(cs)

	if cents%25 == 0 {
		return 25
	}
	return 0
}

func numItemsPoints(is []item) int {
	return (len(is) / 2) * 5
}

func mustTrimmedLengthMultOfThreePoints(itemDescription, price string) int {
	if !prx.MatchString(price) { // prx compiled in `models.go`
		panic(fmt.Errorf("mustTrimmedLengthMultOfThreePoints requires a value formatted as money"))
	}

	d := strings.TrimSpace(itemDescription)
	p, _ := strconv.ParseFloat(price, 64)

	if len(d)%3 == 0 {
		p = p * 0.2
		return int(math.Ceil(p))
	}
	return 0
}

func oddDayPoints(t time.Time) int {
	day := t.Day()
	if day%2 == 1 {
		return 6
	}
	return 0
}

func timeBetweenTwoAndFourPoints(t time.Time) int {
	hour := t.Hour()
	min := t.Minute()

	if (hour == 14 && min > 0) || hour == 15 {
		return 10
	}
	return 0
}
