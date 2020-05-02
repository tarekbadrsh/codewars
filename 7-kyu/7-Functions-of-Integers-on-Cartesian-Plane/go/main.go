package main

import "fmt"

func main() {

	fmt.Println(min(3, 6))
	fmt.Println(SuMin(6))
	fmt.Println(SuMin(100))
	fmt.Println(SuMax(6))
	fmt.Println(SumSum(6))
}

// to Compute sum of digits in all numbers from 1 to n =>  n*(n+1)/2;
func computeN(n int64) int64 {
	return n * (n + 1) / 2
}

/*
min
	1 	1 	1 	1 	1 	1
	1 	2 	2 	2 	2 	2
	1 	2 	3 	3 	3 	3
	1 	2 	3 	4 	4 	4
	1 	2 	3 	4 	5 	5
	1 	2 	3 	4 	5 	6
*/

func min(m int64, max int64) int64 {
	return computeN(m) + ((max - m) * m)
}

// SuMin :
func SuMin(m int) int64 {
	var result int64 = 0
	for i := 1; i <= m; i++ {
		result += min(int64(i), int64(m))
	}
	return result
}

/*
max
	1 	2 	3 	4 	5 	6
	2 	2 	3 	4 	5 	6
	3 	3 	3 	4 	5 	6
	4 	4 	4 	4 	5 	6
	5 	5 	5 	5 	5 	6
	6 	6 	6 	6 	6 	6
*/

func max(m int64, max int64, computeMax int64) int64 {
	result := computeMax - computeN(m)
	result += m * m
	return result
}

// SuMax :
func SuMax(m int) int64 {
	var result int64 = 0
	computeMax := computeN(int64(m))
	for i := 1; i <= m; i++ {
		result += max(int64(i), int64(m), computeMax)
	}
	return result
}

// SumSum :
func SumSum(m int) int64 {
	var result int64 = 0
	computeMax := computeN(int64(m))
	for i := 1; i <= m; i++ {
		result += min(int64(i), int64(m))
		result += max(int64(i), int64(m), computeMax)
	}
	return result
}

///================ best practices ==================///

// SuMinB :
func SuMinB(m int) int64 {
	var n = int64(m)
	return n * (n + 1) * (2*n + 1) / 6
}

// SuMaxB :
func SuMaxB(m int) int64 {
	var n = int64(m)
	return n * (n + 1) * (4*n - 1) / 6
}

// SumSumB :
func SumSumB(m int) int64 {
	var n = int64(m)
	return n * n * (n + 1)
}
