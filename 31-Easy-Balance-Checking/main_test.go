package main

import (
	"reflect"
	"testing"
)

func TestBalance(t *testing.T) {
	var tt = []struct {
		input    string
		expected string
	}{
		{`1000.00!=

125 Market !=:125.45
126 Hardware =34.95
127 Video! 7.45
128 Book :14.32
129 Gasoline ::16.10
		`,
			"Original Balance: 1000.00\n125 Market 125.45 Balance 874.55\n126 Hardware 34.95 Balance 839.60\n127 Video 7.45 Balance 832.15\n128 Book 14.32 Balance 817.83\n129 Gasoline 16.10 Balance 801.73\nTotal expense  198.27\nAverage expense  39.65"},
		{
			`1233.00
125 Hardware;! 24.8?;
123 Flowers 93.5
127 Meat 120.90
120 Picture 34.00
124 Gasoline 11.00
123 Photos;! 71.4?;
122 Picture 93.5
132 Tyres;! 19.00,?;
129 Stamps 13.6
129 Fruits{} 17.6
129 Market;! 128.00?;
121 Gasoline;! 13.6?;
`,
			"Original Balance: 1233.00\n125 Hardware 24.80 Balance 1208.20\n123 Flowers 93.50 Balance 1114.70\n127 Meat 120.90 Balance 993.80\n120 Picture 34.00 Balance 959.80\n124 Gasoline 11.00 Balance 948.80\n123 Photos 71.40 Balance 877.40\n122 Picture 93.50 Balance 783.90\n132 Tyres 19.00 Balance 764.90\n129 Stamps 13.60 Balance 751.30\n129 Fruits 17.60 Balance 733.70\n129 Market 128.00 Balance 605.70\n121 Gasoline 13.60 Balance 592.10\nTotal expense  640.90\nAverage expense  53.41",
		},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := Balance(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("B", func(t *testing.T) {
			res := BalanceB(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("C", func(t *testing.T) {
			res := BalanceC(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkBalance(b *testing.B) {
	input := `1233.00
	125 Hardware;! 24.8?;
	123 Flowers 93.5
	127 Meat 120.90
	120 Picture 34.00
	124 Gasoline 11.00
	123 Photos;! 71.4?;
	122 Picture 93.5
	132 Tyres;! 19.00,?;
	129 Stamps 13.6
	129 Fruits{} 17.6
	129 Market;! 128.00?;
	121 Gasoline;! 13.6?;
	`
	for index := 0; index < b.N; index++ {
		Balance(input)
	}
}
func BenchmarkBalanceB(b *testing.B) {
	input := `1233.00
	125 Hardware;! 24.8?;
	123 Flowers 93.5
	127 Meat 120.90
	120 Picture 34.00
	124 Gasoline 11.00
	123 Photos;! 71.4?;
	122 Picture 93.5
	132 Tyres;! 19.00,?;
	129 Stamps 13.6
	129 Fruits{} 17.6
	129 Market;! 128.00?;
	121 Gasoline;! 13.6?;
	`
	for index := 0; index < b.N; index++ {
		BalanceB(input)
	}
}
func BenchmarkBalanceC(b *testing.B) {
	input := `1233.00
	125 Hardware;! 24.8?;
	123 Flowers 93.5
	127 Meat 120.90
	120 Picture 34.00
	124 Gasoline 11.00
	123 Photos;! 71.4?;
	122 Picture 93.5
	132 Tyres;! 19.00,?;
	129 Stamps 13.6
	129 Fruits{} 17.6
	129 Market;! 128.00?;
	121 Gasoline;! 13.6?;
	`
	for index := 0; index < b.N; index++ {
		BalanceC(input)
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
