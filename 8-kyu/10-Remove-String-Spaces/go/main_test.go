package main

import (
	"reflect"
	"testing"
)

// go test -v
func TestOrderedCount(t *testing.T) {
	var tt = []struct {
		input1   string
		expected string
	}{

		{"8 j 8   mBliB8g  imjB8B8  jl  B", "8j8mBliB8gimjB8B8jlB"},
		{"8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd", "88Bifk8hB8BB8BBBB888chl8BhBfd"},
		{"8aaaaa dddd r     ", "8aaaaaddddr"},
		{"jfBm  gk lf8hg  88lbe8 ", "jfBmgklf8hg88lbe8"},
		{"8j aam", "8jaam"},
	}

	for _, tc := range tt {
		t.Run("NoSpace", func(t *testing.T) {
			res := NoSpace(tc.input1)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
	for _, tc := range tt {
		t.Run("NoSpaceA", func(t *testing.T) {
			res := NoSpaceA(tc.input1)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
	for _, tc := range tt {
		t.Run("NoSpaceB", func(t *testing.T) {
			res := NoSpaceB(tc.input1)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
	for _, tc := range tt {
		t.Run("NoSpaceC", func(t *testing.T) {
			res := NoSpaceC(tc.input1)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkNoSpace(b *testing.B) {
	for index := 0; index < b.N; index++ {
		NoSpace("8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd")
	}
}

func BenchmarkNoSpaceA(b *testing.B) {
	for index := 0; index < b.N; index++ {
		NoSpaceA("8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd")
	}
}

func BenchmarkNoSpaceB(b *testing.B) {
	for index := 0; index < b.N; index++ {
		NoSpaceB("8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd")
	}
}

func BenchmarkNoSpaceC(b *testing.B) {
	for index := 0; index < b.N; index++ {
		NoSpaceC("8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd")
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/8-kyu/10-Remove-String-Spaces/go
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkNoSpaceA-8      5701862               205.7 ns/op            32 B/op          1 allocs/op
BenchmarkNoSpaceB-8      3905565               327.5 ns/op           256 B/op          2 allocs/op
BenchmarkNoSpaceC-8      5937855               176.5 ns/op            56 B/op          3 allocs/op
PASS
ok      github.com/tarekbadrshalaan/codewars/8-kyu/10-Remove-String-Spaces/go   3.541s
*/
