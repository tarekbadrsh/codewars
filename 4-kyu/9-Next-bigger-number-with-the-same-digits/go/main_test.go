package main

import (
	"reflect"
	"testing"
)

// go test -v
func TestOrderedCount(t *testing.T) {
	var tt = []struct {
		input1   int
		expected int
	}{
		{12, 21},
		{513, 531},
		{2017, 2071},
		{414, 441},
		{144, 414},
		{8, -1},
		{123456789, 123456798},
		{1234567890, 1234567908},
		{9876543210, -1},
		{9999999999, -1},
		{59884848459853, 59884848483559},
		{44444, -1},
	}

	for _, tc := range tt {
		t.Run("NextBigger", func(t *testing.T) {
			res := NextBigger(tc.input1)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
	for _, tc := range tt {
		t.Run("NextBiggerA", func(t *testing.T) {
			res := NextBiggerA(tc.input1)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkNextBigger(b *testing.B) {
	for index := 0; index < b.N; index++ {
		NextBigger(59884848459853)
	}
}

func BenchmarkNextBiggerA(b *testing.B) {
	for index := 0; index < b.N; index++ {
		NextBiggerA(59884848459853)
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/4-kyu/9-Next-bigger-number-with-the-same-digits/go
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkNextBigger-8                776           1460090 ns/op          379312 B/op      23707 allocs/op
BenchmarkNextBiggerA-8           6079429               184.7 ns/op            40 B/op          2 allocs/op
PASS
ok      github.com/tarekbadrshalaan/codewars/4-kyu/9-Next-bigger-number-with-the-same-digits/go 2.619s

*/
