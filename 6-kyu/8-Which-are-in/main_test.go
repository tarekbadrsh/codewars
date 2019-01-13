package main

import (
	"reflect"
	"testing"
)

func TestInArray(t *testing.T) {
	var tt = []struct {
		a        []string
		b        []string
		expected []string
	}{
		{[]string{"live", "arp", "strong"}, []string{"lively", "alive", "harp", "sharp", "armstrong"}, []string{"arp", "live", "strong"}},
		{[]string{"live", "arp", "strong"}, []string{"XYZ", "ABC", "HDG"}, []string{}},
	}

	for _, tc := range tt {
		t.Run("", func(t *testing.T) {
			res := InArray(tc.a, tc.b)
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("expected %v; got %v", tc.expected, res)
			}
		})
	}
}
