package main

import (
	"reflect"
	"testing"
)

func TestGap(t *testing.T) {
	var tt = []struct {
		input    []int
		expected int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{-2, -1, -3, -4, -1, -2, -1, -5, -4}, 0},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := MaximumSubarraySum(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := MaximumSubarraySumB(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkMaximumSubarraySum(b *testing.B) {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	for index := 0; index < b.N; index++ {
		MaximumSubarraySum(input)
	}
}

func BenchmarkMaximumSubarraySumB(b *testing.B) {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	for index := 0; index < b.N; index++ {
		MaximumSubarraySumB(input)
	}
}

/*

goos: linux
goarch: amd64

BenchmarkMaximumSubarraySum-8           20000000                52.2 ns/op             0 B/op          0 allocs/op
BenchmarkMaximumSubarraySumB-8          100000000               13.1 ns/op             0 B/op          0 allocs/op

*/
