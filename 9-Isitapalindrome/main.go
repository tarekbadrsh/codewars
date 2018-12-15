package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(IsPalindrome("assssDDDDssssa"))
	fmt.Println(IsPalindrome("2assssDDDDscsssa2"))
}

// IsPalindrome :
func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	l := len(str)
	for i := 0; i < l/2; i++ {
		if str[i] != str[l-i-1] {
			return false
		}
	}
	return true
}
