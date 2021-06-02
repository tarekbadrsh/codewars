package main

import (
	"reflect"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	var tt = []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := LongestCommonPrefix(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := LongestCommonPrefix2(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkLongestCommonPrefix1(b *testing.B) {
	input := []string{"flower", "flow", "flight"}
	for index := 0; index < b.N; index++ {
		LongestCommonPrefix(input)
	}
}
func BenchmarkLongestCommonPrefix2(b *testing.B) {
	input := []string{"flower", "flow", "flight"}
	for index := 0; index < b.N; index++ {
		LongestCommonPrefix2(input)
	}
}

/*

goos: windows
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/__leetcode.com/Easy/1-14-Longest-Common-Prefix/go
cpu: Intel(R) Xeon(R) CPU           W3565  @ 3.20GHz
BenchmarkLongestCommonPrefix1-8          6780216               157.0 ns/op            12 B/op          3 allocs/op
BenchmarkLongestCommonPrefix2-8         13711965                90.77 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/tarekbadrshalaan/codewars/__leetcode.com/Easy/1-14-Longest-Common-Prefix/go  3.661s

*/

//!-bench
