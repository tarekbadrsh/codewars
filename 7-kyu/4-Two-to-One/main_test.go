package main

import "testing"

func TestTwoToOne(t *testing.T) {

	var tt = []struct {
		id       string
		s1       string
		s2       string
		expected string
	}{
		{"1", "abccdd", "hgklmmmnnoo", "abcdghklmno"},
		{"2", "aretheyhere", "yestheyarehere", "aehrsty"},
		{"3", "loopingisfunbutdangerous", "lessdangerousthancoding", "abcdefghilnoprstu"},
	}

	for _, tc := range tt {
		t.Run(tc.id, func(t *testing.T) {
			res := TwoToOne(tc.s1, tc.s2)
			if res != tc.expected {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}
