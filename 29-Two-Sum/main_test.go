package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tt = []struct {
		input    []int
		input2   int
		expected [2]int
	}{
		{[]int{1, 2, 3}, 4, [2]int{0, 2}},
		{[]int{1234, 5678, 9012}, 14690, [2]int{1, 2}},
		{[]int{2, 2, 3}, 4, [2]int{0, 1}},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := TwoSum(tc.input, tc.input2)
			if (res[0] + res[1]) != (tc.expected[0] + tc.expected[1]) {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}
