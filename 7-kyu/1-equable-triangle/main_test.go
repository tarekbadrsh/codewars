package main

import (
	"fmt"
	"testing"
)

func TestEquableTriangle(t *testing.T) {
	var tt = []struct {
		id       string
		a, b, c  int
		expected bool
	}{
		{"1", 5, 12, 13, true},
		{"2", 2, 3, 4, false},
	}

	for _, tc := range tt {
		t.Run(tc.id, func(t *testing.T) {
			res := EquableTriangle(tc.a, tc.b, tc.c)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

	// test EquableTriangle BestPractices
	for _, tc := range tt {
		t.Run(fmt.Sprintf("BestPractices-%v", tc.id), func(t *testing.T) {
			res := EquableTriangleBestPractices(tc.a, tc.b, tc.c)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

}
