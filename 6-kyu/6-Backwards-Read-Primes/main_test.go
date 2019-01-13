package main

import (
	"math/big"
	"testing"
)

func TestBackwardsPrime(t *testing.T) {
	var tt = []struct {
		id       string
		start    int
		stop     int
		expected []int
	}{
		{"1", 1, 100, []int{13, 17, 31, 37, 71, 73, 79, 97}},
		{"2", 1, 31, []int{13, 17, 31}},
		{"3", 501, 599, nil},
	}

	for _, tc := range tt {
		t.Run(tc.id, func(t *testing.T) {
			res := BackwardsPrime(tc.start, tc.stop)
			for i := 0; i < len(tc.expected); i++ {
				if res[i] != tc.expected[i] {
					t.Errorf("expected %v; got %v", tc.expected, res)
				}
			}

		})
	}
}

//!+bench
//go test -v  -bench=.
func BenchmarkIsPrime(b *testing.B) {
	for index := 0; index < b.N; index++ {
		isPrime(index)
	}
}

func BenchmarkProbablyPrime(b *testing.B) {
	for index := 0; index < b.N; index++ {
		big.NewInt(int64(index)).ProbablyPrime(20)
	}
}

/*
	BenchmarkIsPrime-24              3000000               617 ns/op
	BenchmarkProbablyPrime-24         200000              6493 ns/op
*/
//!-bench
