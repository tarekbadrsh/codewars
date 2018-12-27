package main

import "fmt"

func main() {
	fmt.Println(TwoSum([]int{1, 2, 3}, 4))
}

// TwoSum :
func TwoSum(numbers []int, target int) [2]int {
	m := map[int]int{}
	res := [2]int{}
	for k, v := range numbers {
		c := target - v
		if i, ok := m[c]; ok {
			res[0] = k
			res[1] = i
		} else {
			m[v] = k
		}
	}
	return res
}
