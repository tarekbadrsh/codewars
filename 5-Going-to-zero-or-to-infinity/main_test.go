package main

import (
	"fmt"
	"testing"
)

func TestGoing(t *testing.T) {
	var tt = []struct {
		id       string
		input    int
		expected float64
	}{
		{"1", 5, 1.275},
		{"2", 6, 1.2125},
		{"3", 7, 1.173214},
		{"4", 8, 1.146651},
		{"5", 20, 1.052786},
		{"6", 30, 1.034525},
		{"7", 50, 1.020416},
		{"8", 113, 1.008929},
		{"9", 200, 1.005025},
		{"10", 523, 1.001915},
		{"11", 1011, 1.00099},
		{"12", 10110, 1.000098},
	}

	for _, tc := range tt {
		t.Run(tc.id, func(t *testing.T) {
			res := Going(tc.input)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

	// test Going BestPractices
	for _, tc := range tt {
		t.Run(fmt.Sprintf("BestPractices-%v", tc.id), func(t *testing.T) {
			res := GoingBestPractices(tc.input)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}
