package main

import (
	"fmt"
)

func main() {
	fmt.Println(TwoToOne("xyaabbbccccdefww", "xxxxyyyyabklmopq"))
}

// TwoToOne :
func TwoToOne(s1 string, s2 string) string {
	letters := []rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}
	s := s1 + s2
	res := ""
	for _, vl := range letters {
		for _, vs := range s {
			if vl == vs {
				res += string(vl)
				break
			}
		}
	}
	return res
}
