package main

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	var tt = []struct {
		input    string
		expected []int
	}{
		{"ooo", []int{0, 0, 0}},
		{"ioioio", []int{1, 2, 3}},
		{"idoiido", []int{0, 1}},
		{"isoisoiso", []int{1, 4, 25}},
		{"codewars", []int{0}},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := Parse(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}
