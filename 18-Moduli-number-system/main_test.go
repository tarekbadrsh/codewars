package main

import (
	"testing"
)

func TestFromNbToStr(t *testing.T) {
	var tt = []struct {
		n        int64
		sys      []int64
		expected string
	}{
		{779, []int64{8, 7, 5, 3}, "-3--2--4--2-"},
		{187, []int64{8, 7, 5, 3}, "-3--5--2--1-"},
		{15, []int64{8, 6, 5}, "Not applicable"},
		{3450, []int64{17, 5, 3}, "Not applicable"},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := FromNbToStr(tc.n, tc.sys)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

	// test FromNbToStr  BestPractices
	for _, tc := range tt {
		t.Run("BestPractices", func(t *testing.T) {
			res := FromNbToStrB(tc.n, tc.sys)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem

var arr = []int64{97, 31, 17, 13, 11, 7, 5, 3, 2}

func BenchmarkMultiply(b *testing.B) {
	for index := 0; index < b.N; index++ {
		res := Multiply(arr)
		if res != 1535103570 {
			b.Errorf("expected %v; got %v", 1535103570, res)
		}
	}
}

func BenchmarkMultiplyloop(b *testing.B) {
	for index := 0; index < b.N; index++ {
		res := Multiplyloop(arr)
		if res != 1535103570 {
			b.Errorf("expected %v; got %v", 1535103570, res)
		}
	}
}

func BenchmarkGcd(b *testing.B) {
	for index := 0; index < b.N; index++ {
		res := Gcd(17, 5)
		if res != 1 {
			b.Errorf("expected %v; got %v", 1, res)
		}
	}
}

func BenchmarkGcdB(b *testing.B) {
	for index := 0; index < b.N; index++ {
		res := GcdB(17, 5)
		if res != 1 {
			b.Errorf("expected %v; got %v", 1, res)
		}
	}
}

func BenchmarkFromNbToStr(b *testing.B) {
	for index := 0; index < b.N; index++ {
		res := FromNbToStr(779, []int64{8, 7, 5, 3})
		if res != "-3--2--4--2-" {
			b.Errorf("expected %v; got %v", "-3--2--4--2-", res)
		}
	}
}

func BenchmarkFromNbToStrB(b *testing.B) {
	for index := 0; index < b.N; index++ {
		res := FromNbToStrB(779, []int64{8, 7, 5, 3})
		if res != "-3--2--4--2-" {
			b.Errorf("expected %v; got %v", "-3--2--4--2-", res)
		}
	}
}

/*

$ go test -bench . -benchmem
goos: linux
goarch: amd64
BenchmarkMultiply-24            30000000                48.5 ns/op             0 B/op          0 allocs/op
BenchmarkMultiplyloop-24        100000000               12.2 ns/op             0 B/op          0 allocs/op
BenchmarkGcd-24                 50000000                27.2 ns/op             0 B/op          0 allocs/op
BenchmarkGcdB-24                30000000                49.8 ns/op             0 B/op          0 allocs/op
BenchmarkFromNbToStr-24          1000000              1414 ns/op             112 B/op         12 allocs/op
BenchmarkFromNbToStrB-24         2000000               742 ns/op             128 B/op          2 allocs/op

*/
//!-bench
