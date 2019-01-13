package main

import (
	"fmt"
	"testing"
)

func TestNbYear(t *testing.T) {
	var tt = []struct {
		id       string
		p0       int
		percent  float64
		aug      int
		p        int
		expected int
	}{
		{"1", 1500, 5, 100, 5000, 15},
		{"2", 1500000, 2.5, 10000, 2000000, 10},
		{"3", 1500000, 0.25, 1000, 2000000, 94},
		{"4", 1500000, 0.25, -1000, 2000000, 151},
	}

	for _, tc := range tt {
		t.Run(tc.id, func(t *testing.T) {
			res := NbYear(tc.p0, tc.percent, tc.aug, tc.p)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

	// test ToNato BestPractices
	for _, tc := range tt {
		t.Run(fmt.Sprintf("BestPractices-%v", tc.id), func(t *testing.T) {
			res := NbYearBestPractices(tc.p0, tc.percent, tc.aug, tc.p)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}
