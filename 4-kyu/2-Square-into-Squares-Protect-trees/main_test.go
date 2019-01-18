package main

import (
	"reflect"
	"testing"
)

func TestDecompose(t *testing.T) {
	var tt = []struct {
		input    int64
		expected []int64
	}{
		{12, []int64{1, 2, 3, 7, 9}},
		{50, []int64{1, 3, 5, 8, 49}},
		{44, []int64{2, 3, 5, 7, 43}},
		{625, []int64{2, 5, 8, 34, 624}},
		{5, []int64{3, 4}},
		{7100, []int64{2, 3, 5, 119, 7099}},
		{123456, []int64{1, 2, 7, 29, 496, 123455}},
		{1234567, []int64{2, 8, 32, 1571, 1234566}},
		{7654321, []int64{6, 10, 69, 3912, 7654320}},
		{7654322, []int64{1, 4, 11, 69, 3912, 7654321}},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := Decompose(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := DecomposeB(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
}

func BenchmarkDecompose(b *testing.B) {
	var input int64 = 123456
	for index := 0; index < b.N; index++ {
		Decompose(input)
	}
}

func BenchmarkDecomposeB(b *testing.B) {
	var input int64 = 123456
	for index := 0; index < b.N; index++ {
		DecomposeB(input)
	}
}

/*
go test -bench . -benchmem
goos: linux
goarch: amd64
BenchmarkDecompose-8               10000            103008 ns/op             120 B/op          4 allocs/op
BenchmarkDecomposeB-8               5000            363944 ns/op             120 B/op          4 allocs/op

*/
