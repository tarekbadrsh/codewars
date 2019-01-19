package main

import (
	"reflect"
	"testing"
)

func TestValidBraces(t *testing.T) {
	var tt = []struct {
		input    string
		expected bool
	}{
		{"(){}[]", true},
		{"([{}])", true},
		{"(}", false},
		{"[(])", false},
		{"[({)](]", false},
		{"(({[]{()}}[{(())}[]]))", true},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := ValidBraces(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := ValidBracesB(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := ValidBracesC(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkValidBraces(b *testing.B) {
	input := "(({[]{()}}[{(())}[]]))"
	for index := 0; index < b.N; index++ {
		ValidBraces(input)
	}
}

func BenchmarkValidBracesB(b *testing.B) {
	input := "(({[]{()}}[{(())}[]]))"
	for index := 0; index < b.N; index++ {
		ValidBracesB(input)
	}
}

func BenchmarkValidBracesC(b *testing.B) {
	input := "(({[]{()}}[{(())}[]]))"
	for index := 0; index < b.N; index++ {
		ValidBracesC(input)
	}
}

/*

goos: linux
goarch: amd64
BenchmarkValidBraces-8           1000000              1201 ns/op             176 B/op         14 allocs/op
BenchmarkValidBracesB-8          2000000               943 ns/op             240 B/op          4 allocs/op


you can find that Compile regular expression out side 'ValidBracesC' func made a major change in the performance.

@ func with Compile regular expression
BenchmarkValidBracesC-8           200000              6264 ns/op             240 B/op         19 allocs/op

@ func without Compile regular expression
BenchmarkValidBracesC-8           100000             22491 ns/op           40240 B/op         50 allocs/op

*/
