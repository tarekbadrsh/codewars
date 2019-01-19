package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(firstNonRepeating("streSS"))
	fmt.Println(firstNonRepeating("a"))
	fmt.Println(firstNonRepeating("moonmen"))
}

func firstNonRepeating(str string) string {
	repeated := map[rune]int{}
	lowerstr := strings.ToLower(str)
	for _, v := range lowerstr {
		repeated[v]++
	}

	for k, v := range lowerstr {
		if repeated[v] == 1 {
			return string(str[k])
		}
	}
	return ""
}

func firstNonRepeatingB(str string) string {
	for _, c := range str {
		if strings.Count(strings.ToLower(str), strings.ToLower(string(c))) < 2 {
			return string(c)
		}
	}
	return ""
}
