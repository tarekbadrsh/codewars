package main

import (
	"fmt"
	"testing"
)

func TestSumEvenFibonacci(t *testing.T) {
	var tt = []struct {
		id       string
		input    int
		expected int
	}{
		{"1", 8, 10},
		{"2", 111111, 60696},
		{"3", 8675309, 4613732},
	}

	for _, tc := range tt {
		t.Run(tc.id, func(t *testing.T) {
			res := SumEvenFibonacci(tc.input)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

	// test SumEvenFibonacci BestPractices
	for _, tc := range tt {
		t.Run(fmt.Sprintf("BestPractices-%v", tc.id), func(t *testing.T) {
			res := SumEvenFibonacciBestPractices(tc.input)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}

//!+bench
//go test -v  -bench=.
func BenchmarkSumEvenFibonacci1(b *testing.B) {
	for index := 0; index < b.N; index++ {
		res := SumEvenFibonacci(8)
		if res != 10 {
			b.Errorf("expected %v; got %v", 10, res)
		}
	}
}

func BenchmarkSumEvenFibonacci2(b *testing.B) {
	for index := 0; index < b.N; index++ {
		res := SumEvenFibonacci(8675309)
		if res != 4613732 {
			b.Errorf("expected %v; got %v", 4613732, res)
		}
	}
}

func BenchmarkSumEvenFibonacciBestPractices(b *testing.B) {
	for index := 0; index < b.N; index++ {
		res := SumEvenFibonacciBestPractices(8675309)
		if res != 4613732 {
			b.Errorf("expected %v; got %v", 4613732, res)
		}

	}
}

//!-bench
