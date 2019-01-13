package main

import (
	"reflect"
	"testing"
)

func TestValidateSolution(t *testing.T) {
	var tt = []struct {
		input    [][]int
		expected bool
	}{
		{[][]int{
			{5, 3, 4, 6, 7, 8, 9, 1, 2},
			{6, 7, 2, 1, 9, 5, 3, 4, 8},
			{1, 9, 8, 3, 4, 2, 5, 6, 7},
			{8, 5, 9, 7, 6, 1, 4, 2, 3},
			{4, 2, 6, 8, 5, 3, 7, 9, 1},
			{7, 1, 3, 9, 2, 4, 8, 5, 6},
			{9, 6, 1, 5, 3, 7, 2, 8, 4},
			{2, 8, 7, 4, 1, 9, 6, 3, 5},
			{3, 4, 5, 2, 8, 6, 1, 7, 9},
		}, true}, {[][]int{
			{5, 3, 4, 6, 7, 8, 9, 1, 2},
			{6, 7, 2, 1, 9, 0, 3, 4, 8},
			{1, 0, 0, 3, 4, 2, 5, 6, 0},
			{8, 5, 9, 7, 6, 1, 0, 2, 0},
			{4, 2, 6, 8, 5, 3, 7, 9, 1},
			{7, 1, 3, 9, 2, 4, 8, 5, 6},
			{9, 0, 1, 5, 3, 7, 2, 1, 4},
			{2, 8, 7, 4, 1, 9, 6, 3, 5},
			{3, 0, 0, 4, 8, 1, 1, 7, 9},
		}, false}, {[][]int{
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
		}, false}, {[][]int{
			{1, 2, 6, 3, 4, 7, 5, 9, 8},
			{7, 3, 5, 8, 1, 9, 6, 4, 2},
			{1, 9, 4, 2, 7, 5, 8, 6, 3},
			{3, 1, 7, 5, 8, 4, 2, 6, 9},
			{7, 5, 9, 1, 6, 2, 4, 3, 8},
			{4, 8, 2, 9, 3, 6, 7, 1, 5},
			{1, 4, 8, 7, 5, 9, 3, 2, 6},
			{5, 6, 1, 4, 2, 3, 9, 8, 7},
			{2, 7, 3, 6, 9, 1, 8, 5, 4},
		}, false},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := ValidateSolution(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("B", func(t *testing.T) {
			res := ValidateSolutionB(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkValidateSolution(b *testing.B) {
	input := [][]int{
		{1, 2, 6, 3, 4, 7, 5, 9, 8},
		{7, 3, 5, 8, 1, 9, 6, 4, 2},
		{1, 9, 4, 2, 7, 5, 8, 6, 3},
		{3, 1, 7, 5, 8, 4, 2, 6, 9},
		{7, 5, 9, 1, 6, 2, 4, 3, 8},
		{4, 8, 2, 9, 3, 6, 7, 1, 5},
		{1, 4, 8, 7, 5, 9, 3, 2, 6},
		{5, 6, 1, 4, 2, 3, 9, 8, 7},
		{2, 7, 3, 6, 9, 1, 8, 5, 4},
	}
	for index := 0; index < b.N; index++ {
		ValidateSolution(input)
	}
}

func BenchmarkValidateSolutionB(b *testing.B) {
	input := [][]int{
		{1, 2, 6, 3, 4, 7, 5, 9, 8},
		{7, 3, 5, 8, 1, 9, 6, 4, 2},
		{1, 9, 4, 2, 7, 5, 8, 6, 3},
		{3, 1, 7, 5, 8, 4, 2, 6, 9},
		{7, 5, 9, 1, 6, 2, 4, 3, 8},
		{4, 8, 2, 9, 3, 6, 7, 1, 5},
		{1, 4, 8, 7, 5, 9, 3, 2, 6},
		{5, 6, 1, 4, 2, 3, 9, 8, 7},
		{2, 7, 3, 6, 9, 1, 8, 5, 4},
	}
	for index := 0; index < b.N; index++ {
		ValidateSolutionB(input)
	}
}

/*
go test -bench . -benchmem
goos: linux
goarch: amd64
BenchmarkValidateSolution-8      5000000               361 ns/op             248 B/op          5 allocs/op
BenchmarkValidateSolutionB-8    100000000               22.1 ns/op             0 B/op          0 allocs/op

*/
