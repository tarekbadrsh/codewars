package main

import (
	"reflect"
	"testing"
)

func TestEncryptThis(t *testing.T) {
	var tt = []struct {
		input1   string
		input2   string
		expected string
	}{
		{
			"123 Main Street St. Louisville OH 43071,432 Main Long Road St. Louisville OH 43071,786 High Street Pollocksville NY 56432",
			"OH 43071",
			"OH 43071:Main Street St. Louisville,Main Long Road St. Louisville/123,432",
		},
		{
			`123 Main Street St. Louisville OH 43071, 432 Main Long Road St. Louisville OH 43071,786 High Street Pollocksville NY 56432,
			54 Holy Grail Street Niagara Town ZP 32908, 3200 Main Rd. Bern AE 56210,1 Gordon St. Atlanta RE 13000,
			10 Pussy Cat Rd. Chicago EX 34342, 10 Gordon St. Atlanta RE 13000, 58 Gordon Road Atlanta RE 13000,
			22 Tokyo Av. Tedmondville SW 43098, 674 Paris bd. Abbeville AA 45521, 10 Surta Alley Goodtown GG 30654,
			45 Holy Grail Al. Niagara Town ZP 32908, 320 Main Al. Bern AE 56210, 14 Gordon Park Atlanta RE 13000,
			100 Pussy Cat Rd. Chicago EX 34342, 2 Gordon St. Atlanta RE 13000, 5 Gordon Road Atlanta RE 13000,
			2200 Tokyo Av. Tedmondville SW 43098, 67 Paris St. Abbeville AA 45521, 11 Surta Avenue Goodtown GG 30654,
			45 Holy Grail Al. Niagara Town ZP 32918, 320 Main Al. Bern AE 56215, 14 Gordon Park Atlanta RE 13200,
			100 Pussy Cat Rd. Chicago EX 34345, 2 Gordon St. Atlanta RE 13222, 5 Gordon Road Atlanta RE 13001,
			2200 Tokyo Av. Tedmondville SW 43198, 67 Paris St. Abbeville AA 45522, 11 Surta Avenue Goodville GG 30655,
			2222 Tokyo Av. Tedmondville SW 43198, 670 Paris St. Abbeville AA 45522, 114 Surta Avenue Goodville GG 30655,
			2 Holy Grail Street Niagara Town ZP 32908, 3 Main Rd. Bern AE 56210, 77 Gordon St. Atlanta RE 13000,
			100 Pussy Cat Rd. Chicago OH 13201`,
			"AA 45522",
			"AA 45522:Paris St. Abbeville,Paris St. Abbeville/67,670",
		},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := Travel1(tc.input1, tc.input2)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := Travel2(tc.input1, tc.input2)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkTravelThis1(b *testing.B) {
	input := "123 Main Street St. Louisville OH 43071,432 Main Long Road St. Louisville OH 43071,786 High Street Pollocksville NY 56432"
	for index := 0; index < b.N; index++ {
		Travel1(input, "OH 43071")
	}
}

func BenchmarkTravelThis2(b *testing.B) {
	input := "123 Main Street St. Louisville OH 43071,432 Main Long Road St. Louisville OH 43071,786 High Street Pollocksville NY 56432"
	for index := 0; index < b.N; index++ {
		Travel2(input, "OH 43071")
	}
}

/*
goos: windows
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/6-kyu/29-Salesmans-Travel/go
cpu: Intel(R) Xeon(R) CPU           W3565  @ 3.20GHz
BenchmarkTravelThis1-8            308756              3786 ns/op             864 B/op         25 allocs/op
BenchmarkTravelThis2-8            518280              2769 ns/op             720 B/op         16 allocs/op
PASS
ok      github.com/tarekbadrshalaan/codewars/6-kyu/29-Salesmans-Travel/go       3.759s
*/
