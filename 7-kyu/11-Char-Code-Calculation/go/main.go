package main

import (
	"fmt"
)

func main() {
	fmt.Println(GetMiddle("test"))
}

// GetMiddle :
func GetMiddle(s string) string {
	sLength := len(s)
	if sLength%2 == 1 {
		return string(s[sLength/2])
	}
	return string(s[sLength/2-1 : sLength/2+1])
}

///================ other practices ==================///

// GetMiddleA :
func GetMiddleA(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	return s[(n-1)/2 : n/2+1]
}
