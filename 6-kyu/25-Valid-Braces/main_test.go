package main

import (
	"reflect"
	"testing"
)

func TestValidBraces(t *testing.T) {
	var tt = []struct {
		input    string
		expected bool
	}{
		{"(){}[]", true},
		{"([{}])", true},
		{"(}", false},
		{"[(])", false},
		{"[({)](]", false},
		{"(({[]{()}}[{(())}[]]))", true},
	}
	// (({[]{()}}[{(())}[]]))

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := ValidBraces(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
}
