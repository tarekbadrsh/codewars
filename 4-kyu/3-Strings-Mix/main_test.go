package main

import (
	"reflect"
	"testing"
)

func TestMix(t *testing.T) {
	var tt = []struct {
		input1   string
		input2   string
		expected string
	}{

		{"my&friend&Paul has heavy hats! &", "my friend John has many many friends &", "2:nnnnn/1:aaaa/1:hhh/2:mmm/2:yyy/2:dd/2:ff/2:ii/2:rr/=:ee/=:ss"},
		{"mmmmm m nnnnn y&friend&Paul has heavy hats! &", "my frie n d Joh n has ma n y ma n y frie n ds n&", "1:mmmmmm/=:nnnnnn/1:aaaa/1:hhh/2:yyy/2:dd/2:ff/2:ii/2:rr/=:ee/=:ss"},
		{"Are the kids at home? aaaaa fffff", "Yes they are here! aaaaa fffff", "=:aaaaaa/2:eeeee/=:fffff/1:tt/2:rr/=:hh"},
		{"Are they here", "yes, they are here", "2:eeeee/2:yy/=:hh/=:rr"},
		{"uuuuuu", "uuuuuu", "=:uuuuuu"},
		{"looping is fun but dangerous", "less dangerous than coding", "1:ooo/1:uuu/2:sss/=:nnn/1:ii/2:aa/2:dd/2:ee/=:gg"},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := Mix(tc.input1, tc.input2)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := MixB(tc.input1, tc.input2)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkMix(b *testing.B) {
	input1 := "looping is fun but dangerous"
	input2 := "less dangerous than coding"
	for index := 0; index < b.N; index++ {
		Mix(input1, input2)
	}
}

func BenchmarkMixB(b *testing.B) {
	input1 := "looping is fun but dangerous"
	input2 := "less dangerous than coding"
	for index := 0; index < b.N; index++ {
		MixB(input1, input2)
	}
}

/*
go test -bench . -benchmem
goos: linux
goarch: amd64
BenchmarkMix-8            200000             12637 ns/op            4701 B/op         60 allocs/op
BenchmarkMixB-8           500000              3078 ns/op             744 B/op         36 allocs/op
*/
