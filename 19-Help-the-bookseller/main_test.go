package main

import "testing"

func TestStockList(t *testing.T) {
	var tt = []struct {
		a        []string
		b        []string
		expected string
	}{
		{[]string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}, []string{"A", "B", "C", "D"}, "(A : 0) - (B : 1290) - (C : 515) - (D : 600)"},
		{[]string{"ABAR 200", "CDXE 500", "BKWR 250", "BTSQ 890", "DRTY 600"}, []string{"A", "B"}, "(A : 200) - (B : 1140)"},
		{[]string{}, []string{"A", "B", "C", "D"}, ""},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := StockList(tc.a, tc.b)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

	// test StockList other solution
	for _, tc := range tt {
		t.Run("other-", func(t *testing.T) {
			res := StockListB(tc.a, tc.b)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem

func BenchmarkStockList(b *testing.B) {
	x := []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	y := []string{"A", "B", "C", "D"}
	expected := "(A : 0) - (B : 1290) - (C : 515) - (D : 600)"
	for index := 0; index < b.N; index++ {
		res := StockList(x, y)
		if res != expected {
			b.Errorf("expected %v; got %v", expected, res)
		}
	}
}

func BenchmarkStockListB(b *testing.B) {
	x := []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	y := []string{"A", "B", "C", "D"}
	expected := "(A : 0) - (B : 1290) - (C : 515) - (D : 600)"
	for index := 0; index < b.N; index++ {
		res := StockListB(x, y)
		if res != expected {
			b.Errorf("expected %v; got %v", expected, res)
		}
	}
}

/*

$ go test -bench . -benchmem
goos: linux
goarch: amd64
BenchmarkStockList-24             500000              2159 ns/op             332 B/op         10 allocs/op
BenchmarkStockListB-24            100000             12473 ns/op             833 B/op         46 allocs/op

*/
//!-bench
