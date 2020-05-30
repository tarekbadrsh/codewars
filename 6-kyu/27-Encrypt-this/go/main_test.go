package main

import (
	"reflect"
	"testing"
)

func TestEncryptThis(t *testing.T) {
	var tt = []struct {
		input    string
		expected string
	}{
		{"Hello", "72olle"},
		{"good", "103doo"},
		{"hello world", "104olle 119drlo"},
		{"", ""},
		{"A wise old owl lived in an oak", "65 119esi 111dl 111lw 108dvei 105n 97n 111ka"},
		{"The more he saw the less he spoke", "84eh 109ero 104e 115wa 116eh 108sse 104e 115eokp"},
		{"The less he spoke the more he heard", "84eh 108sse 104e 115eokp 116eh 109ero 104e 104dare"},
		{"Why can we not all be like that wise old bird", "87yh 99na 119e 110to 97ll 98e 108eki 116tah 119esi 111dl 98dri"},
		{"Thank you Piotr for all your help", "84kanh 121uo 80roti 102ro 97ll 121ruo 104ple"},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			res := EncryptThis(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			res := EncryptThis2(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkEncryptThis(b *testing.B) {
	input := "The more he saw the less he spoke"
	for index := 0; index < b.N; index++ {
		EncryptThis(input)
	}
}
func BenchmarkEncryptThis2(b *testing.B) {
	input := "The more he saw the less he spoke"
	for index := 0; index < b.N; index++ {
		EncryptThis2(input)
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/6-kyu/27-Encrypt-this/go
BenchmarkEncryptThis-8    	  456482	      2191 ns/op	     584 B/op	      40 allocs/op
BenchmarkEncryptThis2-8   	 1000000	      1136 ns/op	     176 B/op	      18 allocs/op
PASS
ok  	github.com/tarekbadrshalaan/codewars/6-kyu/27-Encrypt-this/go	2.179s
*/
