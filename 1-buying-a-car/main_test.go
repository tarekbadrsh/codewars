package main

import (
	"fmt"
	"testing"
)

func TestNbMonths(t *testing.T) {
	var tt = []struct {
		id                                           string
		startPriceOld, startPriceNew, savingperMonth int
		percentLossByMonth                           float64
		expected                                     [2]int
	}{
		{"1", 2000, 8000, 1000, 1.5, [2]int{6, 766}},
		{"2", 12000, 8000, 1000, 1.5, [2]int{0, 4000}},
		{"3", 8000, 12000, 500, 1.0, [2]int{8, 597}},
		{"4", 18000, 32000, 1500, 1.25, [2]int{8, 332}},
		{"5", 7500, 32000, 300, 1.55, [2]int{25, 122}},
	}
	for _, tc := range tt {
		t.Run(tc.id, func(t *testing.T) {
			res := NbMonths(tc.startPriceOld, tc.startPriceNew, tc.savingperMonth, tc.percentLossByMonth)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

	// test NbMonths BestPractices
	for _, tc := range tt {
		t.Run(fmt.Sprintf("BestPractices-%v", tc.id), func(t *testing.T) {
			res := NbMonthsBestPractices(tc.startPriceOld, tc.startPriceNew, tc.savingperMonth, tc.percentLossByMonth)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}
