package main

import "fmt"

func main() {
	fmt.Println(MaximumSubarraySumB([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

// MaximumSubarraySum :
func MaximumSubarraySum(numbers []int) int {
	max := 0
	for k := range numbers {
		sum := 0
		for _, v := range numbers[k:] {
			sum += v
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

// MaximumSubarraySumB :
func MaximumSubarraySumB(numbers []int) int {
	res, sum := 0, 0
	for _, v := range numbers {
		sum += v
		if res < sum {
			res = sum
		}

		if sum < 0 {
			sum = 0
		}
	}
	return res
}
