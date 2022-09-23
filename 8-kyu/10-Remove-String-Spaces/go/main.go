package main

import (
	"strings"
)

func main() {
}

// NoSpace :
func NoSpace(word string) string {
	result := ""
	for _, v := range word {
		if v != ' ' {
			result += string(v)
		}
	}
	return result
}

///================ other practices ==================///

func NoSpaceA(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

func NoSpaceB(word string) string {
	arr := strings.Fields(word)
	return strings.Join(arr, "")
}

func NoSpaceC(word string) string {
	var b strings.Builder
	for _, v := range word {
		if v != ' ' {
			b.WriteRune(v)
		}
	}
	return b.String()
}
