package main

import (
	"reflect"
	"testing"
)

// go test -v
func TestOrderedCount(t *testing.T) {
	var tt = []struct {
		input1   string
		expected string
	}{
		{"test", "es"},
		{"testing", "t"},
		{"middle", "dd"},
		{"A", "A"},
	}

	for _, tc := range tt {
		t.Run("GetMiddle", func(t *testing.T) {
			res := GetMiddle(tc.input1)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
	for _, tc := range tt {
		t.Run("GetMiddleA", func(t *testing.T) {
			res := GetMiddleA(tc.input1)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkSolution(b *testing.B) {
	for index := 0; index < b.N; index++ {
		GetMiddle("middle")
	}
}

func BenchmarkSolutionA(b *testing.B) {
	for index := 0; index < b.N; index++ {
		GetMiddleA("middle")
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/7-kyu/11-Char-Code-Calculation/go
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkSolution-8     1000000000               0.2579 ns/op          0 B/op          0 allocs/op
BenchmarkSolutionA-8    1000000000               0.2614 ns/op          0 B/op          0 allocs/op
PASS
ok      github.com/tarekbadrshalaan/codewars/7-kyu/11-Char-Code-Calculation/go  0.587s
*/
