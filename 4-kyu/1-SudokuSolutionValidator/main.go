package main

import "fmt"

func main() {

	var testTrue = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	fmt.Println(ValidateSolution(testTrue))
	var testFalse = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 0, 3, 4, 8},
		{1, 0, 0, 3, 4, 2, 5, 6, 0},
		{8, 5, 9, 7, 6, 1, 0, 2, 0},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 0, 1, 5, 3, 7, 2, 1, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 0, 0, 4, 8, 1, 1, 7, 9},
	}

	fmt.Println(ValidateSolution(testFalse))
}

// ValidateSolution :
func ValidateSolution(m [][]int) bool {

	// validate horizontal
	for _, v := range m {
		if sumarr(v) != 45 {
			return false
		}
	}

	// validate vertical
	for i := 0; i < 9; i++ {
		arr := []int{}
		for _, v := range m {
			arr = append(arr, v[i])
		}
		if sumarr(arr) != 45 {
			return false
		}
	}

	// validate square
	from, to := 0, 3
	for h := 0; h < 3; h++ {
		sum := 0
		for v := 0; v < 9; v++ {
			if v%3 == 0 {
				if sum == 0 || sum == 45 {
					sum = 0
				} else {
					return false
				}
			}
			sum += sumarr(m[v][from:to])

		}
		from += 3
		to += 3
	}

	return true
}

func sumarr(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

/**


**/

// ValidateSolutionB :
func ValidateSolutionB(m [][]int) bool {
	var row, col int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			row += m[i][j]
			col += m[j][i]
		}
		if row != 45 || col != 45 {
			return false
		}
		row, col = 0, 0
	}
	if checkThrees(m) {
		return true
	}
	return false
}

func checkThrees(m [][]int) bool {
	var sum int
	var l, n int
	for i := 0; i < 7; i += 3 {
		for j := 0; j < 7; j += 3 {
			for l = i; l < i+3; l++ {
				for n = j; n < j+3; n++ {
					sum += m[l][n]
				}
			}
			if sum != 45 {
				return false
			}
			sum = 0
		}
	}
	return true
}
