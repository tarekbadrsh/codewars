package main

import (
	"reflect"
	"testing"
)

func TestSumFracts(t *testing.T) {
	var tt = []struct {
		input    [][]int
		expected string
	}{
		{[][]int{{1, 2}, {1, 3}, {1, 4}}, "13/12"},
		{[][]int{{1, 3}, {5, 3}}, "2"},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := SumFracts(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("B", func(t *testing.T) {
			res := SumFractsB(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkSumFracts(b *testing.B) {
	input := [][]int{{81345, 15786}, {87546, 11111111}, {43216, 255689}}
	for index := 0; index < b.N; index++ {
		SumFracts(input)
	}
}

func BenchmarkSumFractsB(b *testing.B) {
	input := [][]int{{81345, 15786}, {87546, 11111111}, {43216, 255689}}
	for index := 0; index < b.N; index++ {
		SumFractsB(input)
	}
}

/*

$ go test -bench . -benchmem
goos: linux
goarch: amd64
BenchmarkSumFracts-8             1000000              1791 ns/op              96 B/op          5 allocs/op
BenchmarkSumFractsB-8             300000              4690 ns/op            2224 B/op         50 allocs/op
*/
