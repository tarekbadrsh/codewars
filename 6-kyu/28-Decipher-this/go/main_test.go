package main

import (
	"reflect"
	"testing"
)

func TestDecipherThis(t *testing.T) {
	var tt = []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"65", "A"},
		{"84eh", "The"},
		{"65 66 67", "A B C"},
		{"65 119esi 111dl 111lw 108dvei 105n 97n 111ka", "A wise old owl lived in an oak"},
		{"84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp", "The more he saw the less he spoke"},
		{"84eh 108sse 104e 115eokp 116eh 109ero 104e 104dare", "The less he spoke the more he heard"},
		{"87yh 99na 119e 110to 97ll 98e 108eki 116tah 119esi 111dl 98dri", "Why can we not all be like that wise old bird"},
		{"84kanh 121uo 80roti 102ro 97ll 121ruo 104ple", "Thank you Piotr for all your help"},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			res := DecipherThis(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkDecipherThis(b *testing.B) {
	input := "84kanh 121uo 80roti 102ro 97ll 121ruo 104ple"
	for index := 0; index < b.N; index++ {
		DecipherThis(input)
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/6-kyu/28-Decipher-this/go
BenchmarkDecipherThis-8   	  531801	      2019 ns/op	     512 B/op	      37 allocs/op
PASS
ok  	github.com/tarekbadrshalaan/codewars/6-kyu/28-Decipher-this/go	1.099s
*/
