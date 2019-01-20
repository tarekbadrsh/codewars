package main

import (
	"reflect"
	"testing"
)

func TestGap(t *testing.T) {
	var tt = []struct {
		input    string
		expected string
	}{
		{"a", "a"},
		{"stress", "t"},
		{"moonmen", "e"},
		{"abba", ""},
		{"aa", ""},
		{"~><#~><", "#"},
		{"hello world, eh?", "w"},
		{"sTreSS", "T"},
		{"Go hang a salami, I'm a lasagna hog!", ","},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := firstNonRepeating(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := firstNonRepeatingB(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkFirstNonRepeating(b *testing.B) {
	input := "Go hang a salami, I'm a lasagna hog!"
	for index := 0; index < b.N; index++ {
		firstNonRepeating(input)
	}
}

func BenchmarkFirstNonRepeatingB(b *testing.B) {
	input := "Go hang a salami, I'm a lasagna hog!"
	for index := 0; index < b.N; index++ {
		firstNonRepeatingB(input)
	}
}

/*

goos: linux
goarch: amd64
BenchmarkFirstNonRepeating-8             1000000              2103 ns/op             369 B/op          4 allocs/op
BenchmarkFirstNonRepeatingB-8            1000000              2885 ns/op            1633 B/op         35 allocs/op

*/
