package main

import (
	"fmt"
	"testing"
)

func TestStringConstructing(t *testing.T) {
	var tt = []struct {
		id       string
		a, s     string
		expected int
	}{
		{"1", "aba", "abbabba", 9},
		{"2", "aba", "abaa", 4},
		{"3", "aba", "a", 3},
		{"4", "a", "a", 1},
		{"5", "a", "aaa", 3},
		{"6", "abcdefgh", "hgfedcba", 64},
	}

	for _, tc := range tt {
		t.Run(tc.id, func(t *testing.T) {
			res := StringConstructing(tc.a, tc.s)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

	// test StringConstructing BestPractices
	for _, tc := range tt {
		t.Run(fmt.Sprintf("BestPractices-%v", tc.id), func(t *testing.T) {
			res := StringConstructingBestPractices(tc.a, tc.s)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}
