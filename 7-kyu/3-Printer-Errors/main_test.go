package main

import (
	"fmt"
	"testing"
)

func TestPrinterError(t *testing.T) {
	var tt = []struct {
		id       string
		input    string
		expected string
	}{
		{"1", "aaabbbbhaijjjm", "0/14"},
		{"2", "aaaxbbbbyyhwawiwjjjwwm", "8/22"},
		{"3", "aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz", "3/56"},
		{"4", "kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz", "6/60"},
		{"5", "kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu", "11/65"},
	}

	for _, tc := range tt {
		t.Run(tc.id, func(t *testing.T) {
			res := PrinterError(tc.input)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

	// test PrinterError BestPractices
	for _, tc := range tt {
		t.Run(fmt.Sprintf("BestPractices-%v", tc.id), func(t *testing.T) {
			res := PrinterErrorBestPractices(tc.input)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}
