package main

import (
	"reflect"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var tt = []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"bbbb", 1},
		{"abcabcbb", 3},
		{"pwwkew", 3},
		{"aab", 2},
		{"dvdfd", 3},
		{"dwavdf", 5},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := LengthOfLongestSubstring(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := LengthOfLongestSubstring2(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := LengthOfLongestSubstring3(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkLengthOfLongestSubstring1(b *testing.B) {
	input := "dwavdfzasedfrfdseerrttyhhhhgggfl;kljhgfdsaqwertyhnhju"
	for index := 0; index < b.N; index++ {
		LengthOfLongestSubstring(input)
	}
}

func BenchmarkLengthOfLongestSubstring2(b *testing.B) {
	input := "dwavdfzasedfrfdseerrttyhhhhgggfl;kljhgfdsaqwertyhnhju"
	for index := 0; index < b.N; index++ {
		LengthOfLongestSubstring2(input)
	}
}

func BenchmarkLengthOfLongestSubstring3(b *testing.B) {
	input := "dwavdfzasedfrfdseerrttyhhhhgggfl;kljhgfdsaqwertyhnhju"
	for index := 0; index < b.N; index++ {
		LengthOfLongestSubstring3(input)
	}
}

/*
goos: windows
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/__leetcode.com/Medium/1-3-Longest-Substring-Without-Repeating-Characters/go
cpu: Intel(R) Xeon(R) CPU           W3565  @ 3.20GHz
BenchmarkLengthOfLongestSubstring1-8      108451              9839 ns/op             960 B/op        169 allocs/op
BenchmarkLengthOfLongestSubstring2-8     3549427               316.6 ns/op             0 B/op          0 allocs/op
BenchmarkLengthOfLongestSubstring3-8      222568              4729 ns/op             742 B/op          3 allocs/op
PASS
ok      github.com/tarekbadrshalaan/codewars/__leetcode.com/Medium/1-3-Longest-Substring-Without-Repeating-Characters/go        3.869s

*/

//!-bench
