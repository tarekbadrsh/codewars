package main

import (
	"reflect"
	"testing"
)

func TestSquaresInRect(t *testing.T) {
	var tt = []struct {
		lng      int
		wdth     int
		expected []int
	}{
		{5, 3, []int{3, 2, 1, 1}},
		{3, 5, []int{3, 2, 1, 1}},
		{5, 5, nil},
		{20, 14, []int{14, 6, 6, 2, 2, 2}},
		{14, 20, []int{14, 6, 6, 2, 2, 2}},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := SquaresInRect(tc.lng, tc.wdth)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}
