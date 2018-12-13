package main

import (
	"fmt"
	"testing"
)

func TestEquableTriangle(t *testing.T) {
	var tt = []struct {
		id       string
		roman    string
		expected int
	}{
		{"1", "", 0},
		{"2", "IV", 4},
		{"3", "XXI", 21},
		{"4", "I", 1},
		{"5", "MMVIII", 2008},
		{"6", "MDCLXVI", 1666},
	}

	for _, tc := range tt {
		t.Run(tc.id, func(t *testing.T) {
			res := Decode(tc.roman)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

	// test Decode BestPractices
	for _, tc := range tt {
		t.Run(fmt.Sprintf("BestPractices-%v", tc.id), func(t *testing.T) {
			res := DecodeBestPractices(tc.roman)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

}
