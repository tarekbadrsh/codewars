package main

import (
	"reflect"
	"testing"
)

func TestMeeting(t *testing.T) {
	var tt = []struct {
		s        string
		expected string
	}{
		{"Victoria:Schwarz", "(SCHWARZ, VICTORIA)"},
		{"Alexis:Wahl;John:Bell;Victoria:Schwarz", "(BELL, JOHN)(SCHWARZ, VICTORIA)(WAHL, ALEXIS)"},
		{"Alexis:Wahl;John:Bell;Victoria:Schwarz;Abba:Dorny;Grace:Meta;Ann:Arno;Madison:STAN;Alex:Cornwell;Lewis:Kern;Megan:Stan;Alex:Korn",
			"(ARNO, ANN)(BELL, JOHN)(CORNWELL, ALEX)(DORNY, ABBA)(KERN, LEWIS)(KORN, ALEX)(META, GRACE)(SCHWARZ, VICTORIA)(STAN, MADISON)(STAN, MEGAN)(WAHL, ALEXIS)"},
		{"John:Gates;Michael:Wahl;Megan:Bell;Paul:Dorries;James:Dorny;Lewis:Steve;Alex:Meta;Elizabeth:Russel;Anna:Korn;Ann:Kern;Amber:Cornwell",
			"(BELL, MEGAN)(CORNWELL, AMBER)(DORNY, JAMES)(DORRIES, PAUL)(GATES, JOHN)(KERN, ANN)(KORN, ANNA)(META, ALEX)(RUSSEL, ELIZABETH)(STEVE, LEWIS)(WAHL, MICHAEL)"},
	}
	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := Meeting(tc.s)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

	// solution B
	for _, tc := range tt {
		t.Run("B", func(t *testing.T) {
			res := MeetingB(tc.s)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

	// solution C
	for _, tc := range tt {
		t.Run("C", func(t *testing.T) {
			res := MeetingC(tc.s)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

	// solution D
	for _, tc := range tt {
		t.Run("D", func(t *testing.T) {
			res := MeetingD(tc.s)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem

func BenchmarkMeetingA(b *testing.B) {
	input := "Alexis:Wahl;John:Bell;Victoria:Schwarz;Abba:Dorny;Grace:Meta;Ann:Arno;Madison:STAN;Alex:Cornwell;Lewis:Kern;Megan:Stan;Alex:Korn"
	for index := 0; index < b.N; index++ {
		Meeting(input)
	}
}

func BenchmarkMeetingB(b *testing.B) {
	input := "Alexis:Wahl;John:Bell;Victoria:Schwarz;Abba:Dorny;Grace:Meta;Ann:Arno;Madison:STAN;Alex:Cornwell;Lewis:Kern;Megan:Stan;Alex:Korn"
	for index := 0; index < b.N; index++ {
		MeetingB(input)
	}
}

func BenchmarkMeetingC(b *testing.B) {
	input := "Alexis:Wahl;John:Bell;Victoria:Schwarz;Abba:Dorny;Grace:Meta;Ann:Arno;Madison:STAN;Alex:Cornwell;Lewis:Kern;Megan:Stan;Alex:Korn"
	for index := 0; index < b.N; index++ {
		MeetingC(input)
	}
}

func BenchmarkMeetingD(b *testing.B) {
	input := "Alexis:Wahl;John:Bell;Victoria:Schwarz;Abba:Dorny;Grace:Meta;Ann:Arno;Madison:STAN;Alex:Cornwell;Lewis:Kern;Megan:Stan;Alex:Korn"
	for index := 0; index < b.N; index++ {
		MeetingD(input)
	}
}

/*

$ go test -bench . -benchmem
goos: linux
goarch: amd64
BenchmarkMeetingA-24              200000              8079 ns/op            2224 B/op         31 allocs/op
BenchmarkMeetingB-24              200000              7368 ns/op            1856 B/op         32 allocs/op
BenchmarkMeetingC-24              300000              6572 ns/op            1376 B/op         22 allocs/op
BenchmarkMeetingD-24              200000              7420 ns/op            1824 B/op         33 allocs/op

*/
