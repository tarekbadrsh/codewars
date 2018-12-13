package main

import (
	"fmt"
	"testing"
)

func TestDuplicateCount(t *testing.T) {
	var tt = []struct {
		id       string
		input    string
		expected int
	}{
		{"1", "", 0},
		{"2", "abcde", 0},
		{"3", "abcdea", 1},
		{"4", "abcdeaB11", 3},
		{"5", "indivisibility", 1},
	}

	for _, tc := range tt {
		t.Run(tc.id, func(t *testing.T) {
			res := duplicateCount(tc.input)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

	// test duplicateCount BestPractices
	for _, tc := range tt {
		t.Run(fmt.Sprintf("BestPractices-%v", tc.id), func(t *testing.T) {
			res := duplicateCountBestPractices(tc.input)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

}
