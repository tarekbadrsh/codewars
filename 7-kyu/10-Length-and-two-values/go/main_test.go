package main

import (
	"reflect"
	"testing"
)

// go test -v
func TestOrderedCount(t *testing.T) {
	var tt = []struct {
		n              int
		input1, input2 string
		expected       []string
	}{
		{5, "true", "false", []string{"true", "false", "true", "false", "true"}},
		{20, "blue", "red", []string{"blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red"}},
		{0, "", "", []string{}},
	}

	for _, tc := range tt {
		t.Run("Alternate", func(t *testing.T) {
			res := Alternate(tc.n, tc.input1, tc.input2)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
	for _, tc := range tt {
		t.Run("AlternateA", func(t *testing.T) {
			res := AlternateA(tc.n, tc.input1, tc.input2)
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
		Alternate(20, "blue", "red")
	}
}

func BenchmarkSolutionA(b *testing.B) {
	for index := 0; index < b.N; index++ {
		AlternateA(20, "blue", "red")
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/7-kyu/10-Length-and-two-values/go
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkSolution
BenchmarkSolution-8    	 1942627	       620.5 ns/op
BenchmarkSolutionA
BenchmarkSolutionA-8   	 8397763	       142.2 ns/op
PASS
ok  	github.com/tarekbadrshalaan/codewars/7-kyu/10-Length-and-two-values/go	3.176s
*/
