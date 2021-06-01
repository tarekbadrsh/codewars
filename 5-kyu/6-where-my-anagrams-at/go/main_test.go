package main

import (
	"reflect"
	"testing"
)

// go test -v  -bench=.
func TestDirReduc(t *testing.T) {
	var tt = []struct {
		input1   string
		input2   []string
		expected []string
	}{
		{"abba", []string{"aabb", "abcd", "bbaa", "dada"}, []string{"aabb", "bbaa"}},
		{"racer", []string{"carer", "arcre", "carre", "racrs", "racers", "arceer", "raccer", "carrer", "cerarr"}, []string{"carer", "arcre", "carre"}},
		{"laser", []string{"lazing", "lazy", "lacer"}, nil},
		{"", []string{"Hello", "there"}, nil},
		{"yolo", []string{}, nil},
	}

	for _, tc := range tt {
		t.Run("Anagrams1", func(t *testing.T) {
			res := Anagrams(tc.input1, tc.input2)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
	for _, tc := range tt {
		t.Run("Anagrams2", func(t *testing.T) {
			res := Anagrams2(tc.input1, tc.input2)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
	for _, tc := range tt {
		t.Run("Anagrams3", func(t *testing.T) {
			res := Anagrams3(tc.input1, tc.input2)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkAnagrams(b *testing.B) {
	for index := 0; index < b.N; index++ {
		Anagrams("racer", []string{"carer", "arcre", "carre", "racrs", "racers", "arceer", "raccer", "carrer", "cerarr"})
	}
}
func BenchmarkAnagrams2(b *testing.B) {
	for index := 0; index < b.N; index++ {
		Anagrams2("racer", []string{"carer", "arcre", "carre", "racrs", "racers", "arceer", "raccer", "carrer", "cerarr"})
	}
}
func BenchmarkAnagrams3(b *testing.B) {
	for index := 0; index < b.N; index++ {
		Anagrams3("racer", []string{"carer", "arcre", "carre", "racrs", "racers", "arceer", "raccer", "carrer", "cerarr"})
	}
}

/*
goos: windows
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/5-kyu/6-/go
cpu: Intel(R) Xeon(R) CPU           W3565  @ 3.20GHz
BenchmarkAnagrams-8       668364              1707 ns/op             112 B/op          3 allocs/op
BenchmarkAnagrams2-8       44006             27836 ns/op            5371 B/op        178 allocs/op
BenchmarkAnagrams3-8      244140              4520 ns/op             912 B/op         33 allocs/op
PASS
ok      github.com/tarekbadrshalaan/codewars/5-kyu/6-/go        5.570s
*/
