package main

import (
	"reflect"
	"testing"
)

func TestSuMin(t *testing.T) {
	var tt = []struct {
		vector, circle1, circle2 []float64
		expected                 []float64
	}{
		{[]float64{5, 5}, []float64{10, 20, 42}, []float64{-100, -42, 69}, []float64{-108.21, -66.64}},
		{[]float64{46.0, 58.0}, []float64{0, 0, 100}, []float64{0, 0, 100}, []float64{46.0, 58.0}},
		{[]float64{50.0, 88.0}, []float64{-25.0, 128, 100}, []float64{50, 50, 100}, []float64{125, 10}},
		{[]float64{120, 58.0}, []float64{100, 76, 36}, []float64{120, -62, 50}, []float64{147.78, -87.0}},
		{[]float64{5, 5}, []float64{10, 20, 42}, []float64{-100, -42, 60}, []float64{-107.14, -63.43}},
	}

	for _, tc := range tt {
		t.Run("MapVector", func(t *testing.T) {
			res := MapVector(tc.vector, tc.circle1, tc.circle2)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}

	// for _, tc := range tt {
	// 	t.Run("MapVectorB", func(t *testing.T) {
	// 		res := MapVectorB(tc.vector, tc.circle1, tc.circle2)
	// 		if !reflect.DeepEqual(res, tc.expected) {
	// 			t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
	// 		}
	// 	})
	// }

}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkSuMin(b *testing.B) {
	for index := 0; index < b.N; index++ {
		MapVector([]float64{5, 5}, []float64{10, 20, 42}, []float64{-100, -42, 69})
	}
}

/*
goos: linux
goarch: amd64

pkg: github.com/tarekbadrshalaan/codewars/7-kyu/8-2D-Vector-Mapping/go
BenchmarkSuMin-8   	35414258	        33.3 ns/op	      16 B/op	       1 allocs/op
PASS

ok  	github.com/tarekbadrshalaan/codewars/7-kyu/8-2D-Vector-Mapping/go	2.018s
*/
