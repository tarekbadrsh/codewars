package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solution("45", "1"))
}

// Solution :
func Solution(a, b string) string {
	if len(a) < len(b) {
		return fmt.Sprintf("%s%s%s", a, b, a)
	}
	return fmt.Sprintf("%s%s%s", b, a, b)
}

///================ other practices ==================///

func SolutionA(a, b string) string {
	return min(a, b) + max(a, b) + min(a, b)
}

func min(x, y string) string {
	if len(x) > len(y) {
		return y
	}
	return x
}
func max(x, y string) string {
	if len(x) > len(y) {
		return x
	}
	return y
}
