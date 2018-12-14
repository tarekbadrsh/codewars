package main

import (
	"testing"
)

func TestGap(t *testing.T) {
	var tt = []struct {
		id       string
		g, m, n  int
		expected []int
	}{
		{"1", 2, 100, 110, []int{101, 103}},
		{"2", 4, 100, 110, []int{103, 107}},
		{"3", 6, 100, 110, nil},
		{"4", 8, 300, 400, []int{359, 367}},
		{"5", 10, 300, 400, []int{337, 347}},
	}

	for _, tc := range tt {
		t.Run(tc.id, func(t *testing.T) {
			res := Gap(tc.g, tc.m, tc.n)
			if tc.expected == nil && res != nil {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}

			if tc.expected != nil && (res[0] != tc.expected[0] || res[1] != tc.expected[1]) {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}
