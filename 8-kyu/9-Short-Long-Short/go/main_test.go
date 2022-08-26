package main

import (
	"reflect"
	"testing"
)

// go test -v
func TestOrderedCount(t *testing.T) {
	var tt = []struct {
		input1, input2 string
		expected       string
	}{
		{"45", "1", "1451"},
		{"13", "200", "1320013"},
		{"Soon", "Me", "MeSoonMe"},
		{"U", "False", "UFalseU"},
	}

	for _, tc := range tt {
		t.Run("Solution", func(t *testing.T) {
			res := Solution(tc.input1, tc.input2)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
	for _, tc := range tt {
		t.Run("SolutionA", func(t *testing.T) {
			res := SolutionA(tc.input1, tc.input2)
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
		Solution("Soon", "Me")
	}
}

func BenchmarkSolutionA(b *testing.B) {
	for index := 0; index < b.N; index++ {
		SolutionA("Soon", "Me")
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/8-kyu/9-Short-Long-Short/go
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkSolution
BenchmarkSolution-8    	 6721120	       180.1 ns/op
BenchmarkSolutionA
BenchmarkSolutionA-8   	56342286	        20.06 ns/op
PASS
ok  	github.com/tarekbadrshalaan/codewars/8-kyu/9-Short-Long-Short/go	2.548s
*/
