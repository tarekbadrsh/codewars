package main

import (
	"reflect"
	"testing"
)

func gettestdate() []struct {
	input1   []string
	expected []string
} {
	var tt = []struct {
		input1   []string
		expected []string
	}{
		{[]string{"WEST", "WEST", "EAST", "EAST", "WEST", "SOUTH", "NORTH", "SOUTH"}, []string{"WEST", "SOUTH"}},
		{[]string{"NORTH", "WEST", "EAST", "EAST", "SOUTH"}, []string{"NORTH", "EAST", "SOUTH"}},
		{[]string{"SOUTH", "SOUTH", "SOUTH", "NORTH", "SOUTH", "NORTH", "WEST", "EAST"}, []string{"SOUTH", "SOUTH"}},
		{[]string{"WEST", "EAST", "EAST", "SOUTH", "NORTH"}, []string{"EAST"}},
		{[]string{"WEST", "EAST", "SOUTH", "WEST", "EAST"}, []string{"SOUTH"}},
		{[]string{"NORTH", "EAST", "SOUTH", "NORTH", "SOUTH", "NORTH"}, []string{"NORTH", "EAST"}},
		{[]string{"NORTH", "SOUTH", "WEST", "NORTH", "NORTH"}, []string{"WEST", "NORTH", "NORTH"}},
		{[]string{"NORTH", "NORTH", "NORTH", "NORTH", "SOUTH", "SOUTH", "WEST"}, []string{"NORTH", "NORTH", "WEST"}},
		{[]string{"NORTH", "WEST", "EAST", "WEST", "EAST", "WEST"}, []string{"NORTH", "WEST"}},
		{[]string{"EAST", "WEST", "NORTH", "EAST", "NORTH"}, []string{"NORTH", "EAST", "NORTH"}},
		{[]string{"WEST", "NORTH", "EAST", "NORTH", "WEST"}, []string{"WEST", "NORTH", "EAST", "NORTH", "WEST"}},
		{[]string{"EAST", "SOUTH", "SOUTH", "SOUTH", "EAST", "NORTH", "NORTH"}, []string{"EAST", "SOUTH", "SOUTH", "SOUTH", "EAST", "NORTH", "NORTH"}},
		{[]string{"NORTH", "EAST", "SOUTH", "EAST", "WEST", "NORTH"}, []string{"NORTH", "EAST"}},
		{[]string{"WEST", "EAST", "SOUTH", "EAST", "WEST", "WEST"}, []string{"SOUTH", "WEST"}},
		{[]string{"SOUTH", "NORTH", "SOUTH", "EAST", "SOUTH", "EAST", "EAST"}, []string{"SOUTH", "EAST", "SOUTH", "EAST", "EAST"}},
	}
	return tt
}
func TestDirReduc(t *testing.T) {
	for _, tc := range gettestdate() {
		t.Run("DirReduc", func(t *testing.T) {
			res := DirReduc(tc.input1)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range gettestdate() {
		t.Run("DirReduc2", func(t *testing.T) {
			res := DirReduc2(tc.input1)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range gettestdate() {
		t.Run("DirReduc3", func(t *testing.T) {
			res := DirReduc3(tc.input1)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func benchmardate() []string {
	return []string{"SOUTH", "NORTH", "SOUTH", "EAST", "SOUTH", "EAST", "EAST"}
}
func BenchmarkDirReduc(b *testing.B) {
	for index := 0; index < b.N; index++ {
		DirReduc(benchmardate())
	}
}

func BenchmarkDirReduc2(b *testing.B) {

	for index := 0; index < b.N; index++ {
		DirReduc2(benchmardate())
	}
}

func BenchmarkDirReduc3(b *testing.B) {
	for index := 0; index < b.N; index++ {
		DirReduc3(benchmardate())
	}
}

/*
goos: windows
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/5-kyu/5-Directions-Reduction/go
cpu: Intel(R) Xeon(R) CPU           W3565  @ 3.20GHz
BenchmarkDirReduc-8     27224746                44.11 ns/op            0 B/op          0 allocs/op
BenchmarkDirReduc2-8    18363224                59.27 ns/op            0 B/op          0 allocs/op
BenchmarkDirReduc3-8     1571698               785.3 ns/op           240 B/op          4 allocs/op
PASS
ok      github.com/tarekbadrshalaan/codewars/5-kyu/5-Directions-Reduction/go    6.626s
*/
