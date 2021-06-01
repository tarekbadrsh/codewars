package main

import (
	"reflect"
	"testing"
)

// go test -v  -bench=.
func TestDirReduc(t *testing.T) {
	var tt = []struct {
		input1   string
		expected string
	}{
		{"Hey fellow warriors", "Hey wollef sroirraw"},
		{"Burgers are my favorite fruit", "sregruB are my etirovaf tiurf"},
		{"Pizza is the best vegetable", "azziP is the best elbategev"},
		{"This sentence is a sentence", "This ecnetnes is a ecnetnes"},
	}

	for _, tc := range tt {
		t.Run("SpinWords", func(t *testing.T) {
			res := SpinWords(tc.input1)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
	for _, tc := range tt {
		t.Run("SpinWords2", func(t *testing.T) {
			res := SpinWords2(tc.input1)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkSpinWords(b *testing.B) {
	for index := 0; index < b.N; index++ {
		SpinWords("Hey fellow warriors")
	}
}
func BenchmarkSpinWords2(b *testing.B) {
	for index := 0; index < b.N; index++ {
		SpinWords2("Hey fellow warriors")
	}
}

/*
goos: windows
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/5-kyu/7-Stop-gninnipS-My-sdroW/go
cpu: Intel(R) Xeon(R) CPU           W3565  @ 3.20GHz
BenchmarkSpinWords-8             1000000              1106 ns/op             216 B/op          7 allocs/op
BenchmarkSpinWords2-8             698352              1464 ns/op             144 B/op         16 allocs/op
PASS
ok      github.com/tarekbadrshalaan/codewars/5-kyu/7-Stop-gninnipS-My-sdroW/go  2.278s
*/
