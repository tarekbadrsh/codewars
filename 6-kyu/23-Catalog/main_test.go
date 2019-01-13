package main

import (
	"reflect"
	"testing"
)

func TestCatalog(t *testing.T) {
	var tt = []struct {
		input    string
		input2   string
		expected string
	}{
		{`<prod><name>drill</name><prx>99</prx><qty>5</qty></prod>

<prod><name>hammer</name><prx>10</prx><qty>50</qty></prod>

<prod><name>screwdriver</name><prx>5</prx><qty>51</qty></prod>

<prod><name>table saw</name><prx>1099.99</prx><qty>5</qty></prod>

<prod><name>saw</name><prx>9</prx><qty>10</qty></prod>`, "saw", "table saw > prx: $1099.99 qty: 5\nsaw > prx: $9 qty: 10"},

		{`<prod><name>drill</name><prx>99</prx><qty>5</qty></prod>

<prod><name>hammer</name><prx>10</prx><qty>50</qty></prod>

<prod><name>screwdriver</name><prx>5</prx><qty>51</qty></prod>

<prod><name>table saw</name><prx>1099.99</prx><qty>5</qty></prod>

<prod><name>hammer man </name><prx>9</prx><qty>10</qty></prod>`, "hammer", "hammer > prx: $10 qty: 50\nhammer man  > prx: $9 qty: 10"},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := Catalog(tc.input, tc.input2)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}

	for _, tc := range tt {
		t.Run("B", func(t *testing.T) {
			res := CatalogB(tc.input, tc.input2)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected\n%v\ngot\n%v", tc.expected, res)
			}
		})
	}
}

//!+bench
// go test -v  -bench=.
// go test -bench . -benchmem
func BenchmarkCatalog(b *testing.B) {
	input := `<prod><name>drill</name><prx>99</prx><qty>5</qty></prod>

<prod><name>hammer</name><prx>10</prx><qty>50</qty></prod>

<prod><name>screwdriver</name><prx>5</prx><qty>51</qty></prod>

<prod><name>table saw</name><prx>1099.99</prx><qty>5</qty></prod>

<prod><name>saw</name><prx>9</prx><qty>10</qty></prod>`
	input2 := "saw"
	for index := 0; index < b.N; index++ {
		Catalog(input, input2)
	}
}

func BenchmarkCatalogB(b *testing.B) {
	input := `<prod><name>drill</name><prx>99</prx><qty>5</qty></prod>

<prod><name>hammer</name><prx>10</prx><qty>50</qty></prod>

<prod><name>screwdriver</name><prx>5</prx><qty>51</qty></prod>

<prod><name>table saw</name><prx>1099.99</prx><qty>5</qty></prod>

<prod><name>saw</name><prx>9</prx><qty>10</qty></prod>`
	input2 := "saw"
	for index := 0; index < b.N; index++ {
		CatalogB(input, input2)
	}
}

/*

$ go test -bench . -benchmem
goos: linux
goarch: amd64
BenchmarkCatalog-24                30000             44680 ns/op            7808 B/op        203 allocs/op
BenchmarkCatalogB-24               30000             44740 ns/op            7808 B/op        203 allocs/op

*/
