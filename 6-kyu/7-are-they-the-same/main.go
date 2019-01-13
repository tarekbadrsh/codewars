package main

import "fmt"

func main() {
	a := []int{121, 144, 19, 161, 19, 144, 19, 11}
	b := []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}
	fmt.Println(Comp(a, b))
}

// Comp :
func Comp(array1 []int, array2 []int) bool {
	if len(array1) != len(array2) || array1 == nil {
		return false
	}
	a1 := map[int]int{}
	for _, v := range array1 {
		a1[v*v]++
	}

	for _, v := range array2 {
		if a1[v] < 1 {
			return false
		}
		a1[v]--
	}

	return true
}
