package main

import (
	"reflect"
	"testing"
)

func TestThirt(t *testing.T) {
	var tt = []struct {
		input    int
		expected int
	}{
		{8529, 79},
		{85299258, 31},
		{5634, 57},
		{1111111111, 71},
		{987654321, 30},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := Thirt(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := ThirtB(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkThirt(b *testing.B) {
	input := 987654321
	for index := 0; index < b.N; index++ {
		Thirt(input)
	}
}

func BenchmarkThirtB(b *testing.B) {
	input := 987654321
	for index := 0; index < b.N; index++ {
		ThirtB(input)
	}
}

/*

goos: linux
goarch: amd64
BenchmarkThirt-24        1000000              1346 ns/op              80 B/op          9 allocs/op
BenchmarkThirtB-24      20000000               101 ns/op               0 B/op          0 allocs/op

*/
