package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAlphaneumericCountPoints(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{
			input:    "Target",
			expected: 6,
		},
		{
			input:    "M&M Corner Market",
			expected: 14,
		},
	}

	for _, tc := range testCases {
		actual := alphanumericCountPoints(tc.input)
		assert.Equal(t, tc.expected, actual, fmt.Sprintf("expected %s to receive %d points - got %d", tc.input, tc.expected, actual))
	}
}

func TestMustNoCentsPoints(t *testing.T) {
	testCases := []struct {
		input       string
		expected    int
		shouldPanic bool
	}{
		{
			input:       "35.35",
			expected:    0,
			shouldPanic: false,
		},
		{
			input:       "9.00",
			expected:    50,
			shouldPanic: false,
		},
		{
			input:       "ff.ff",
			expected:    0,
			shouldPanic: true,
		},
	}

	for _, tc := range testCases {
		if tc.shouldPanic {
			assert.Panics(t, func() { _ = mustNoCentsPoints(tc.input) }, fmt.Sprintf("invalid format should panic - %s", tc.input))
		} else {
			actual := mustNoCentsPoints(tc.input)
			assert.Equal(t, tc.expected, actual, fmt.Sprintf("expected %s to receive %d points - got %d", tc.input, tc.expected, actual))
		}
	}
}

func TestMustIsMultipleOfQuarterPoints(t *testing.T) {
	testCases := []struct {
		input       string
		expected    int
		shouldPanic bool
	}{
		{
			input:       "9.00",
			expected:    25,
			shouldPanic: false,
		},
		{
			input:       "35.35",
			expected:    0,
			shouldPanic: false,
		},
		{
			input:       "ff.ff",
			expected:    0,
			shouldPanic: true,
		},
	}

	for _, tc := range testCases {
		if tc.shouldPanic {
			assert.Panics(t, func() { _ = mustIsMultipleOfQuarterPoints(tc.input) }, fmt.Sprintf("invalid format should panic - %s", tc.input))
		} else {
			actual := mustIsMultipleOfQuarterPoints(tc.input)
			assert.Equal(t, tc.expected, actual, fmt.Sprintf("expected %s to receive %d points - got %d", tc.input, tc.expected, actual))
		}
	}
}

func TestNumItemsPoints(t *testing.T) {
	testCases := []struct {
		items    []item
		expected int
	}{
		{
			items: []item{{
				ShortDescription: "Mountain Dew 12PK",
				Price:            "6.49",
			}, {
				ShortDescription: "Emils Cheese Pizza",
				Price:            "12.25",
			}, {
				ShortDescription: "Knorr Creamy Chicken",
				Price:            "1.26",
			}, {
				ShortDescription: "Doritos Nacho Cheese",
				Price:            "3.35",
			}, {
				ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
				Price:            "12.00",
			}},
			expected: 10,
		},
		{
			items:    []item{},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		actual := numItemsPoints(tc.items)
		assert.Equal(t, tc.expected, actual, fmt.Sprintf("expected %d items to receive %d points - got %d", len(tc.items), tc.expected, actual))
	}
}

func TestTrimmedLengthMultOfThreePoints(t *testing.T) {
	testCases := []struct {
		itemDescription string
		price           string
		expected        int
	}{
		{
			itemDescription: "Mountain Dew 12PK",
			price:           "6.49",
			expected:        0,
		},
		{
			itemDescription: "Emils Cheese Pizza",
			price:           "12.25",
			expected:        3,
		},
		{
			itemDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
			price:           "12.00",
			expected:        3,
		},
	}

	for _, tc := range testCases {
		actual := mustTrimmedLengthMultOfThreePoints(tc.itemDescription, tc.price)
		assert.Equal(t, tc.expected, actual, fmt.Sprintf("expected %s to receive %d points - got %d", tc.itemDescription, tc.expected, actual))
	}
}

func TestOddDayPoints(t *testing.T) {
	testCases := []struct {
		date     string
		expected int
	}{
		{
			date:     "2022-01-01",
			expected: 6,
		},
		{
			date:     "2022-03-20",
			expected: 0,
		},
	}

	for _, tc := range testCases {
		time, _ := time.Parse("2006-01-02", tc.date)
		actual := oddDayPoints(time)
		assert.Equal(t, tc.expected, actual, fmt.Sprintf("expected %s to result in %d points - got %d", tc.date, tc.expected, actual))
	}
}

func TestTimeBetweenTwoAndFourPoints(t *testing.T) {
	testCases := []struct {
		time     string
		expected int
	}{
		{
			time:     "14:33",
			expected: 10,
		},
		{
			time:     "13:01",
			expected: 0,
		},
		{
			time:     "14:00",
			expected: 0,
		},
		{
			time:     "16:00",
			expected: 0,
		},
	}

	for _, tc := range testCases {
		time, _ := time.Parse("15:04", tc.time)
		actual := timeBetweenTwoAndFourPoints(time)
		assert.Equal(t, tc.expected, actual, fmt.Sprintf("expected %s to result in %d points - got %d", tc.time, tc.expected, actual))
	}
}
