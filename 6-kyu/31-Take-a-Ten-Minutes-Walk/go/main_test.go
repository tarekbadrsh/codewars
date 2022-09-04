package main

import (
	"reflect"
	"testing"
)

// go test -v
func TestOrderedCount(t *testing.T) {
	var tt = []struct {
		input1   []rune
		expected bool
	}{
		{[]rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}, true},
		{[]rune{'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e'}, false},
		{[]rune{'w'}, false},
		{[]rune{'n', 'n', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}, false},
		{[]rune{'e', 'e', 'e', 'e', 'w', 'w', 's', 's', 's', 's'}, false},
	}

	for _, tc := range tt {
		t.Run("Solution", func(t *testing.T) {
			res := IsValidWalk(tc.input1)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("Solution", func(t *testing.T) {
			res := IsValidWalkA(tc.input1)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkIsValidWalk(b *testing.B) {
	for index := 0; index < b.N; index++ {
		IsValidWalk([]rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'})
	}
}

func BenchmarkIsValidWalkA(b *testing.B) {
	for index := 0; index < b.N; index++ {
		IsValidWalkA([]rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'})
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/6-kyu/31-Take-a-Ten-Minutes-Walk/go
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkIsValidWalk-8    	176310464	        6.886 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsValidWalkA-8   	89172384	        13.13 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/tarekbadrshalaan/codewars/6-kyu/31-Take-a-Ten-Minutes-Walk/go	3.093s
*/
