package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(duplicateCount("indivisibility"))
}

func duplicateCount(s1 string) int {
	lst := map[rune]int{}
	res := 0
	for _, v := range strings.ToLower(s1) {
		lst[v]++
		if lst[v] == 2 {
			res++
		}
	}
	return res
}

func duplicateCountBestPractices(s1 string) int {
	var c int
	h := map[rune]int{}
	for _, r := range strings.ToLower(s1) {
		if h[r]++; h[r] == 2 {
			c++
		}
	}
	return c
}
