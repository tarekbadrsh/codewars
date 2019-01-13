package main

import (
	"math"
	"testing"
)

func assertFuzzy(act float64, exp float64) bool {
	inrange := false
	merr := 1e-9
	e := math.Abs((act - exp) / exp)
	inrange = (e <= merr)
	return inrange
}

func TestPosAverage(t *testing.T) {
	var tt = []struct {
		input    string
		expected float64
	}{
		{"6900690040, 4690606946, 9990494604", 26.666666666666668},
		{"466960, 069060, 494940, 060069, 060090, 640009, 496464, 606900, 004000, 944096", 26.666666666666668},
		{"444996, 699990, 666690, 096904, 600644, 640646, 606469, 409694, 666094, 606490", 29.259259259259256},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := PosAverage(tc.input)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

	//solution B
	for _, tc := range tt {
		t.Run("B", func(t *testing.T) {
			res := PosAverageB(tc.input)
			if !assertFuzzy(res, tc.expected) {
				t.Errorf("Expected should be near: %1.9e\n but got: %1.9e\n", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem

func BenchmarkPosAverage(b *testing.B) {
	input := "466960, 069060, 494940, 060069, 060090, 640009, 496464, 606900, 004000, 944096"
	expected := 26.666666666666668
	for index := 0; index < b.N; index++ {
		res := PosAverage(input)
		if !assertFuzzy(res, expected) {
			b.Errorf("expected %v; got %v", expected, res)
		}
	}
}

func BenchmarkPosAverageB(b *testing.B) {
	input := "466960, 069060, 494940, 060069, 060090, 640009, 496464, 606900, 004000, 944096"
	expected := 26.666666666666668
	for index := 0; index < b.N; index++ {
		res := PosAverageB(input)
		if !assertFuzzy(res, expected) {
			b.Errorf("expected %v; got %v", expected, res)
		}
	}
}

/*

$ go test -bench . -benchmem
goos: linux
goarch: amd64
BenchmarkPosAverage-24           1000000              1513 ns/op             160 B/op          1 allocs/op
BenchmarkPosAverageB-24          1000000              1801 ns/op             160 B/op          1 allocs/op

*/
//!-bench
