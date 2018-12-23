package main

import (
	"reflect"
	"testing"
)

func TestNthRank(t *testing.T) {
	var tt = []struct {
		st       string
		we       []int
		n        int
		expected string
	}{
		{
			st:       "Addison,Jayden,Sofia,Michael,Andrew,Lily,Benjamin",
			we:       []int{4, 2, 1, 4, 3, 1, 2},
			n:        4,
			expected: "Benjamin",
		},
		{
			st:       "Elijah,Chloe,Elizabeth,Matthew,Natalie,Jayden",
			we:       []int{1, 3, 5, 5, 3, 6},
			n:        2,
			expected: "Matthew",
		},
		{
			st:       "Aubrey,Olivai,Abigail,Chloe,Andrew,Elizabeth",
			we:       []int{3, 1, 4, 4, 3, 2},
			n:        4,
			expected: "Abigail",
		},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := NthRank(tc.st, tc.we, tc.n)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

	// solution B
	for _, tc := range tt {
		t.Run("B", func(t *testing.T) {
			res := NthRankB(tc.st, tc.we, tc.n)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}
