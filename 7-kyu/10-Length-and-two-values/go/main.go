package main

import (
	"fmt"
)

func main() {
	fmt.Println(Alternate(5, "true", "false"))
}

// Alternate :

func Alternate(n int, firstValue string, secondValue string) []string {
	res := []string{}
	input := secondValue
	for i := 0; i < n; i++ {
		if input == firstValue {
			input = secondValue
			res = append(res, input)
		} else if input == secondValue {
			input = firstValue
			res = append(res, input)
		}
	}
	return res
}

///================ other practices ==================///

// AlternateA :
func AlternateA[T any](n int, x T, y T) []T {
	res := make([]T, n)
	xy := [2]T{x, y}
	for i := 0; i < n; i++ {
		res[i] = xy[i%2]
	}
	return res
}
