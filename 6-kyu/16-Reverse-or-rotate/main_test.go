package main

import (
	"reflect"
	"testing"
)

func TestRevrot(t *testing.T) {
	var tt = []struct {
		s        string
		n        int
		expected string
	}{
		{"563000655734469485", 4, "0365065073456944"},
		{"1234", 0, ""},
		{"", 0, ""},
		{"1234", 5, ""},
		{"733049910872815764", 5, "330479108928157"},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := Revrot(tc.s, tc.n)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}
