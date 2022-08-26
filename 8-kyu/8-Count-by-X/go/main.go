package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountBy(1, 5))
}

// CountBy :
func CountBy(x, n int) []int {

	res := []int{}
	for i := x; i <= (x * n); i += x {
		res = append(res, i)
	}
	return res
}

///================ other practices ==================///

func CountByA(x, n int) []int {
	arr := []int{}

	for i := 1; i <= n; i++ {
		arr = append(arr, i*x)
	}

	return arr
}
