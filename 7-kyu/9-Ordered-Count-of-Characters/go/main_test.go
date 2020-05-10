package main

import (
	"reflect"
	"testing"
)

func TestOrderedCount(t *testing.T) {
	var tt = []struct {
		input    string
		expected []Tuple
	}{
		{"abracadabra", []Tuple{{'a', 5}, {'b', 2}, {'r', 2}, {'c', 1}, {'d', 1}}},
		{"Code Wars", []Tuple{{'C', 1}, {'o', 1}, {'d', 1}, {'e', 1}, {' ', 1}, {'W', 1}, {'a', 1}, {'r', 1}, {'s', 1}}},
		{"", []Tuple{}},
	}

	for _, tc := range tt {
		t.Run("OrderedCount", func(t *testing.T) {
			res := OrderedCount(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("OrderedCountB", func(t *testing.T) {
			res := OrderedCountA(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("OrderedCountC", func(t *testing.T) {
			res := OrderedCountB(tc.input)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("\nexpected:%v\ngot:%v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkOrderedCount(b *testing.B) {
	for index := 0; index < b.N; index++ {
		OrderedCount("Code Wars")
	}
}

func BenchmarkOrderedCountA(b *testing.B) {
	for index := 0; index < b.N; index++ {
		OrderedCountA("Code Wars")
	}
}

func BenchmarkOrderedCountB(b *testing.B) {
	for index := 0; index < b.N; index++ {
		OrderedCountB("Code Wars")
	}
}

func BenchmarkOrderedCountA_long(b *testing.B) {
	for index := 0; index < b.N; index++ {
		OrderedCountA(`
		Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.
Why do we use it?
It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like).
Where does it come from?
Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of "de Finibus Bonorum et Malorum" (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, "Lorem ipsum dolor sit amet..", comes from a line in section 1.10.32.The standard chunk of Lorem Ipsum used since the 1500s is reproduced below for those interested. Sections 1.10.32 and 1.10.33 from "de Finibus Bonorum et Malorum" by Cicero are also reproduced in their exact original form, accompanied by English versions from the 1914 translation by H. Rackham.
`)
	}
}

func BenchmarkOrderedCountB_long(b *testing.B) {
	for index := 0; index < b.N; index++ {
		OrderedCountB(`
		Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.
Why do we use it?
It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here', making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for 'lorem ipsum' will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like).
Where does it come from?
Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of "de Finibus Bonorum et Malorum" (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, "Lorem ipsum dolor sit amet..", comes from a line in section 1.10.32.The standard chunk of Lorem Ipsum used since the 1500s is reproduced below for those interested. Sections 1.10.32 and 1.10.33 from "de Finibus Bonorum et Malorum" by Cicero are also reproduced in their exact original form, accompanied by English versions from the 1914 translation by H. Rackham.
`)
	}
}

/*
go test -bench . -benchmem
goos: linux
goarch: amd64
pkg: github.com/tarekbadrshalaan/codewars/7-kyu/9-Ordered-Count-of-Characters/go
BenchmarkOrderedCount-8         	 2082000	       560 ns/op	     625 B/op	       2 allocs/op
BenchmarkOrderedCountA-8        	 9616179	       123 ns/op	     448 B/op	       1 allocs/op
BenchmarkOrderedCountB-8        	 8869267	       132 ns/op	     320 B/op	       1 allocs/op
BenchmarkOrderedCountA_long-8   	  142101	      7371 ns/op	    3136 B/op	       3 allocs/op
BenchmarkOrderedCountB_long-8   	  187825	      5645 ns/op	    2240 B/op	       3 allocs/op
PASS
ok  	github.com/tarekbadrshalaan/codewars/7-kyu/9-Ordered-Count-of-Characters/go	6.634s

*/
