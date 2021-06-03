package main

import (
	"reflect"
	"testing"
)

// go test -v  -bench=.
func TestDirReduc(t *testing.T) {
	var tt = []struct {
		input    int
		expected int
	}{
		{10, 22},
		{50, 175},
		{1000, 8488},
		{100, 447},
		{5637, 75711},
	}

	for _, tc := range tt {
		t.Run("DblLinear1", func(t *testing.T) {
			res := DblLinear(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
	for _, tc := range tt {
		t.Run("DblLinear2", func(t *testing.T) {
			res := DblLinear2(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
	for _, tc := range tt {
		t.Run("DblLinear3", func(t *testing.T) {
			res := DblLinear3(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
	for _, tc := range tt {
		t.Run("DblLinear4", func(t *testing.T) {
			res := DblLinear4(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkDblLinear(b *testing.B) {
	for index := 0; index < b.N; index++ {
		DblLinear(5637)
	}
}
func BenchmarkDblLinear2(b *testing.B) {
	for index := 0; index < b.N; index++ {
		DblLinear2(5637)
	}
}
func BenchmarkDblLinear3(b *testing.B) {
	for index := 0; index < b.N; index++ {
		DblLinear3(5637)
	}
}
func BenchmarkDblLinear4(b *testing.B) {
	for index := 0; index < b.N; index++ {
		DblLinear4(5637)
	}
}

/*
goos: windows
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/4-kyu/7-Twice-linear/go
cpu: Intel(R) Xeon(R) CPU           W3565  @ 3.20GHz
BenchmarkDblLinear-8                 277           4139814 ns/op          817901 B/op        367 allocs/op
BenchmarkDblLinear2-8                  1        2799636000 ns/op          951896 B/op       5993 allocs/op
BenchmarkDblLinear3-8              10000            108722 ns/op          214256 B/op         17 allocs/op
BenchmarkDblLinear4-8              12514            121706 ns/op          214256 B/op         17 allocs/op
PASS
ok      github.com/tarekbadrshalaan/codewars/4-kyu/7-Twice-linear/go    10.801s
*/
