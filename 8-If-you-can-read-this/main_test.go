package main

import (
	"fmt"
	"testing"
)

func TestToNato(t *testing.T) {
	var tt = []struct {
		id       string
		input    string
		expected string
	}{
		{"1", "If you can read", "India Foxtrot Yankee Oscar Uniform Charlie Alfa November Romeo Echo Alfa Delta"},
		{"2", "Did not see that coming", "Delta India Delta November Oscar Tango Sierra Echo Echo Tango Hotel Alfa Tango Charlie Oscar Mike India November Golf"},
		{"3", "go for it!", "Golf Oscar Foxtrot Oscar Romeo India Tango !"},
	}
	for _, tc := range tt {
		t.Run(tc.id, func(t *testing.T) {
			res := ToNato(tc.input)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}

	// test ToNato BestPractices
	for _, tc := range tt {
		t.Run(fmt.Sprintf("BestPractices-%v", tc.id), func(t *testing.T) {
			res := ToNatoBestPractices(tc.input)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}
