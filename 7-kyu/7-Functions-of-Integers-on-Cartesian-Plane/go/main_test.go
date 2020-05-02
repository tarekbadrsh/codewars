package main

import (
	"reflect"
	"testing"
)

func TestSuMin(t *testing.T) {
	var tt = []struct {
		input    int
		expected int64
	}{
		{5, 55},
		{100, 338350},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := SuMin(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := SuMinB(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
}

func TestSuMax(t *testing.T) {
	var tt = []struct {
		input    int
		expected int64
	}{
		{5, 95},
		{100, 671650},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := SuMax(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := SuMaxB(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
}

func TestSumSum(t *testing.T) {
	var tt = []struct {
		input    int
		expected int64
	}{
		{5, 150},
		{100, 1010000},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := SumSum(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := SumSumB(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkSuMin(b *testing.B) {
	input := 100
	for index := 0; index < b.N; index++ {
		SuMin(input)
	}
}

func BenchmarkSuMinB(b *testing.B) {
	input := 100
	for index := 0; index < b.N; index++ {
		SuMinB(input)
	}
}
func BenchmarkSuMax(b *testing.B) {
	input := 100
	for index := 0; index < b.N; index++ {
		SuMax(input)
	}
}

func BenchmarkSuMaxB(b *testing.B) {
	input := 100
	for index := 0; index < b.N; index++ {
		SuMaxB(input)
	}
}
func BenchmarkSumSum(b *testing.B) {
	input := 100
	for index := 0; index < b.N; index++ {
		SumSum(input)
	}
}

func BenchmarkSumSumB(b *testing.B) {
	input := 100
	for index := 0; index < b.N; index++ {
		SumSumB(input)
	}
}

/*
goos: linux
goarch: amd64
pkg: func

BenchmarkSuMin-8     	10645449	       114 ns/op	       0 B/op	       0 allocs/op
BenchmarkSuMinB-8    	1000000000	         0.254 ns/op	       0 B/op	       0 allocs/op

BenchmarkSuMax-8     	 9899728	       119 ns/op	       0 B/op	       0 allocs/op
BenchmarkSuMaxB-8    	1000000000	         0.254 ns/op	       0 B/op	       0 allocs/op

BenchmarkSumSum-8    	 7821774	       152 ns/op	       0 B/op	       0 allocs/op
BenchmarkSumSumB-8   	1000000000	         0.254 ns/op	       0 B/op	       0 allocs/op

PASS
ok  	func	5.800s

*/
