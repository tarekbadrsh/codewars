package main

import "testing"

func TestComp(t *testing.T) {
	var tt = []struct {
		a        []int
		b        []int
		expected bool
	}{
		{[]int{121, 144, 19, 161, 19, 144, 19, 11}, []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}, true},
		{[]int{121, 144, 19, 161, 19, 144, 19, 11}, []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}, false},
		{nil, []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}, false},
		{[]int{}, []int{}, true},
		{nil, nil, false},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := Comp(tc.a, tc.b)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}
