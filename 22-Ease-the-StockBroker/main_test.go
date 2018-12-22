package main

import "testing"

func TestBalanceStatement(t *testing.T) {
	var tt = []struct {
		input    string
		expected string
	}{
		{"GOOG 300 542.0 B, AAPL 50 145.0 B, CSCO 250.0 29 B, GOOG 200 580.0 S", "Buy: 169850 Sell: 116000; Badly formed 1: CSCO 250.0 29 B ;"},
		{"ZNGA 1300 2.66 B, CLH15.NYM 50 56.32 B, OWW 1000 11.623 B, OGG 20 580.1 B", "Buy: 29499 Sell: 0"},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := BalanceStatement(tc.input)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}
