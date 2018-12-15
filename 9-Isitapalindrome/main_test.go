package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var tt = []struct {
		id       string
		input    string
		expected bool
	}{
		{"1", "a", true},
		{"2", "aba", true},
		{"3", "Abba", true},
		{"4", "hello", false},
	}
	for _, tc := range tt {
		t.Run(tc.id, func(t *testing.T) {
			res := IsPalindrome(tc.input)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}
