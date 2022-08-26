package main

import (
	"reflect"
	"testing"
)

// go test -v
func TestOrderedCount(t *testing.T) {
	var tt = []struct {
		inputx, inputn int
		expected       []int
	}{
		{1, 5, []int{1, 2, 3, 4, 5}},
		{2, 5, []int{2, 4, 6, 8, 10}},
		{3, 5, []int{3, 6, 9, 12, 15}},
		{50, 5, []int{50, 100, 150, 200, 250}},
		{100, 5, []int{100, 200, 300, 400, 500}},
	}

	for _, tc := range tt {
		t.Run("CountBy", func(t *testing.T) {
			res := CountBy(tc.inputx, tc.inputn)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
	for _, tc := range tt {
		t.Run("CountByA", func(t *testing.T) {
			res := CountByA(tc.inputx, tc.inputn)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkCountBy(b *testing.B) {
	for index := 0; index < b.N; index++ {
		CountBy(100, 5)
	}
}

func BenchmarkCountByA(b *testing.B) {
	for index := 0; index < b.N; index++ {
		CountByA(100, 5)
	}
}

/*
go test -bench . -benchmem
goos: linux
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/8-kyu/8-Count-by-X/go
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkCountBy
BenchmarkCountBy-8    	 9356689	       118.1 ns/op
BenchmarkCountByA
BenchmarkCountByA-8   	10797034	       117.8 ns/op
PASS
ok  	github.com/tarekbadrshalaan/codewars/8-kyu/8-Count-by-X/go	2.625s
*/
