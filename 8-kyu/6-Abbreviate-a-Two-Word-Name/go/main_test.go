package main

import (
	"reflect"
	"testing"
)

func TestAbbrevName(t *testing.T) {
	var tt = []struct {
		input    string
		expected string
	}{
		{"Sam Harris", "S.H"},
		{"Patrick Feenan", "P.F"},
		{"Evan Cole", "E.C"},
		{"P Favuzzi", "P.F"},
		{"David Mendieta", "D.M"},
		{"te be", "T.B"},
		{"aaa sss", "A.S"},
	}

	for _, tc := range tt {
		t.Run("AbbrevName", func(t *testing.T) {
			res := AbbrevName(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("AbbrevNameB", func(t *testing.T) {
			res := AbbrevNameB(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("AbbrevNameC", func(t *testing.T) {
			res := AbbrevNameC(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkSuMin(b *testing.B) {
	input := "Patrick Feenan"
	for index := 0; index < b.N; index++ {
		AbbrevName(input)
	}
}

func BenchmarkSuMinB(b *testing.B) {
	input := "Patrick Feenan"
	for index := 0; index < b.N; index++ {
		AbbrevNameB(input)
	}
}

func BenchmarkSuMinC(b *testing.B) {
	input := "Patrick Feenan"
	for index := 0; index < b.N; index++ {
		AbbrevNameC(input)
	}
}

/*

goos: linux
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/8-kyu/6-Abbreviate-a-Two-Word-Name/go
BenchmarkSuMin-8    	25782319	        45.8 ns/op	      16 B/op	       2 allocs/op
BenchmarkSuMinB-8   	25848057	        44.8 ns/op	       3 B/op	       1 allocs/op
BenchmarkSuMinC-8   	52884985	        22.5 ns/op	       3 B/op	       1 allocs/op
PASS
ok  	github.com/tarekbadrshalaan/codewars/8-kyu/6-Abbreviate-a-Two-Word-Name/go	4.552s

*/
