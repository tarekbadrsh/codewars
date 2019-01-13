package main

import (
	"reflect"
	"testing"
)

func TestDigitalRoot(t *testing.T) {
	var tt = []struct {
		input    int
		expected int
	}{
		{16, 7},
		{195, 6},
		{992, 2},
		{167346, 9},
		{0, 0},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := DigitalRoot(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("B", func(t *testing.T) {
			res := DigitalRootB(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkDigitalRoot(b *testing.B) {
	input := 167346
	for index := 0; index < b.N; index++ {
		DigitalRoot(input)
	}
}

func BenchmarkDigitalRootB(b *testing.B) {
	input := 167346
	for index := 0; index < b.N; index++ {
		DigitalRootB(input)
	}
}

/*

$ go test -bench . -benchmem
goos: linux
goarch: amd64
BenchmarkDigitalRoot-24          2000000               879 ns/op              56 B/op          7 allocs/op
BenchmarkDigitalRootB-24        2000000000               0.75 ns/op            0 B/op          0 allocs/op

*/
