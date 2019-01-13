package main

import "testing"

func TestSolution(t *testing.T) {
	var tt = []struct {
		input    string
		idx      uint
		expected uint
	}{
		{"((1)23(45))(aB)", 0, 10},
		{"((1)23(45))(aB)", 1, 3},
		{"((1)23(45))(aB)", 6, 9},
		{"((>)|?(*'))(yZ)", 11, 14},
		{`(|"M=-H;!(.(l)Mea)(&f<+)C(q)$)OTv8K`, 23, 0},
		{"(W&C()1)o()().", 10, 0},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res, err := Solution(tc.input, tc.idx)
			if tc.expected == 0 && err == nil {
				t.Errorf("expected error; got %v", res)
			}
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}
