package main

import (
	"reflect"
	"testing"
)

func TestByState(t *testing.T) {
	var tt = []struct {
		input    string
		expected string
	}{
		{`John Daggett, 341 King Road, Plymouth MA
Alice Ford, 22 East Broadway, Richmond VA
Sal Carpenter, 73 6th Street, Boston MA`,
			"Massachusetts\n..... John Daggett 341 King Road Plymouth Massachusetts\n..... Sal Carpenter 73 6th Street Boston Massachusetts\n Virginia\n..... Alice Ford 22 East Broadway Richmond Virginia"},
		{
			`John Daggett, 341 King Road, Plymouth MA
Alice Ford, 22 East Broadway, Richmond VA
Orville Thomas, 11345 Oak Bridge Road, Tulsa OK
Terry Kalkas, 402 Lans Road, Beaver Falls PA
Eric Adams, 20 Post Road, Sudbury MA
Hubert Sims, 328A Brook Road, Roanoke VA
Amy Wilde, 334 Bayshore Pkwy, Mountain View CA
Sal Carpenter, 73 6th Street, Boston MA`,
			"California\n..... Amy Wilde 334 Bayshore Pkwy Mountain View California\n Massachusetts\n..... Eric Adams 20 Post Road Sudbury Massachusetts\n..... John Daggett 341 King Road Plymouth Massachusetts\n..... Sal Carpenter 73 6th Street Boston Massachusetts\n Oklahoma\n..... Orville Thomas 11345 Oak Bridge Road Tulsa Oklahoma\n Pennsylvania\n..... Terry Kalkas 402 Lans Road Beaver Falls Pennsylvania\n Virginia\n..... Alice Ford 22 East Broadway Richmond Virginia\n..... Hubert Sims 328A Brook Road Roanoke Virginia",
		},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := ByState(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("B", func(t *testing.T) {
			res := ByStateB(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("C", func(t *testing.T) {
			res := ByStateC(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkByState(b *testing.B) {
	input := `John Daggett, 341 King Road, Plymouth MA
Alice Ford, 22 East Broadway, Richmond VA
Orville Thomas, 11345 Oak Bridge Road, Tulsa OK
Terry Kalkas, 402 Lans Road, Beaver Falls PA
Eric Adams, 20 Post Road, Sudbury MA
Hubert Sims, 328A Brook Road, Roanoke VA
Amy Wilde, 334 Bayshore Pkwy, Mountain View CA
Sal Carpenter, 73 6th Street, Boston MA`
	for index := 0; index < b.N; index++ {
		ByState(input)
	}
}

func BenchmarkByStateB(b *testing.B) {
	input := `John Daggett, 341 King Road, Plymouth MA
Alice Ford, 22 East Broadway, Richmond VA
Orville Thomas, 11345 Oak Bridge Road, Tulsa OK
Terry Kalkas, 402 Lans Road, Beaver Falls PA
Eric Adams, 20 Post Road, Sudbury MA
Hubert Sims, 328A Brook Road, Roanoke VA
Amy Wilde, 334 Bayshore Pkwy, Mountain View CA
Sal Carpenter, 73 6th Street, Boston MA`
	for index := 0; index < b.N; index++ {
		ByStateB(input)
	}
}

func BenchmarkByStateC(b *testing.B) {
	input := `John Daggett, 341 King Road, Plymouth MA
Alice Ford, 22 East Broadway, Richmond VA
Orville Thomas, 11345 Oak Bridge Road, Tulsa OK
Terry Kalkas, 402 Lans Road, Beaver Falls PA
Eric Adams, 20 Post Road, Sudbury MA
Hubert Sims, 328A Brook Road, Roanoke VA
Amy Wilde, 334 Bayshore Pkwy, Mountain View CA
Sal Carpenter, 73 6th Street, Boston MA`
	for index := 0; index < b.N; index++ {
		ByStateC(input)
	}
}

/*

$ go test -bench . -benchmem
goos: linux
goarch: amd64
BenchmarkByState-24               100000             17180 ns/op            5301 B/op         80 allocs/op
BenchmarkByStateB-24              100000             19694 ns/op            5317 B/op         70 allocs/op
BenchmarkByStateC-24               30000             59395 ns/op           46736 B/op         63 allocs/op

*/
