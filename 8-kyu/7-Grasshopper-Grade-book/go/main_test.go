package main

import (
	"reflect"
	"testing"
)

func TestSuMin(t *testing.T) {
	var tt = []struct {
		a, b, c  int
		expected rune
	}{
		{95, 90, 93, 'A'}, {100, 85, 96, 'A'}, {92, 93, 94, 'A'},
		{70, 70, 100, 'B'}, {82, 85, 87, 'B'}, {84, 79, 85, 'B'}, {89, 89, 90, 'B'},
		{70, 70, 70, 'C'}, {75, 70, 79, 'C'}, {60, 82, 76, 'C'},
		{65, 70, 59, 'D'}, {66, 62, 68, 'D'}, {58, 62, 70, 'D'},
		{44, 55, 52, 'F'}, {48, 55, 52, 'F'}, {58, 59, 60, 'F'},
	}

	for _, tc := range tt {
		t.Run("GetGrade", func(t *testing.T) {
			res := GetGrade(tc.a, tc.b, tc.c)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("GetGradeA", func(t *testing.T) {
			res := GetGradeA(tc.a, tc.b, tc.c)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("GetGradeB", func(t *testing.T) {
			res := GetGradeB(tc.a, tc.b, tc.c)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("GetGradeC", func(t *testing.T) {
			res := GetGradeC(tc.a, tc.b, tc.c)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("GetGradeD", func(t *testing.T) {
			res := GetGradeD(tc.a, tc.b, tc.c)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}

}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkGetGrade(b *testing.B) {
	for index := 0; index < b.N; index++ {
		GetGrade(58, 59, 60)
	}
}
func BenchmarkGetGradeA(b *testing.B) {
	for index := 0; index < b.N; index++ {
		GetGradeA(58, 59, 60)
	}
}
func BenchmarkGetGradeB(b *testing.B) {
	for index := 0; index < b.N; index++ {
		GetGradeB(58, 59, 60)
	}
}
func BenchmarkGetGradeC(b *testing.B) {
	for index := 0; index < b.N; index++ {
		GetGradeC(58, 59, 60)
	}
}
func BenchmarkGetGradeD(b *testing.B) {
	for index := 0; index < b.N; index++ {
		GetGradeD(58, 59, 60)
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/8-kyu/7-Grasshopper-Grade-book/go
BenchmarkGetGrade-8    	1000000000	         0.255 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetGradeA-8   	1000000000	         0.252 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetGradeB-8   	1000000000	         0.253 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetGradeC-8   	16319127	        71.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetGradeD-8   	 4268965	       297 ns/op	     236 B/op	       1 allocs/op
PASS
ok  	github.com/tarekbadrshalaan/codewars/8-kyu/7-Grasshopper-Grade-book/go	3.649s
*/
