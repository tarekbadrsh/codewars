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
	fmt.Println(ValidateSolutionC(testTrue))
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
	fmt.Println(ValidateSolutionC(testFalse))
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

///================ other practices ==================///

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
	return checkThrees(m)
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

///================ other practices ==================///
type MainArr struct {
	elem  int
	exist bool
}

func GetMainArr() [10]MainArr {
	return [10]MainArr{{elem: 0, exist: true}, {elem: 1, exist: false}, {elem: 2, exist: false}, {elem: 3, exist: false}, {elem: 4, exist: false}, {elem: 5, exist: false}, {elem: 6, exist: false}, {elem: 7, exist: false}, {elem: 8, exist: false}, {elem: 9, exist: false}}
}

// ValidateSolutionC :
func ValidateSolutionC(m [][]int) bool {

	// validate horizontal
	for _, v := range m {
		mainarr := GetMainArr()
		for _, x := range v {
			if mainarr[x].exist {
				return false
			}
			mainarr[x].exist = true
		}
	}

	// validate vertical
	for i := 0; i < 9; i++ {
		mainarr := GetMainArr()
		for x := 0; x < 9; x++ {
			if mainarr[m[x][i]].exist {
				return false
			}
			mainarr[m[x][i]].exist = true
		}
	}

	// validate square
	/*
		               |----|----|----| ----|----|----| ----|----|----|
			square1	   |0X0Y|0X1Y|0X2Y| 1X0Y|1X1Y|1X2Y| 2X0Y|2X1Y|2X2Y|
		               |----|----|----| ----|----|----| ----|----|----|
			square2	   |0X3Y|0X4Y|0X5Y| 1X3Y|1X4Y|1X5Y| 2X3Y|2X4Y|2X5Y|
		               |----|----|----| ----|----|----| ----|----|----|
			square3	   |0X6Y|0X7Y|0X8Y| 1X6Y|1X7Y|1X8Y| 2X6Y|2X7Y|2X8Y|
		               |----|----|----| ----|----|----| ----|----|----|
			square4	   |3X0Y|3X1Y|3X2Y| 4X0Y|4X1Y|4X2Y| 5X0Y|5X1Y|5X2Y|
		               |----|----|----| ----|----|----| ----|----|----|
			square5	   |3X3Y|3X4Y|3X5Y| 4X3Y|4X4Y|4X5Y| 5X3Y|5X4Y|5X5Y|
		               |----|----|----| ----|----|----| ----|----|----|
			square6	   |3X6Y|3X7Y|3X8Y| 4X6Y|4X7Y|4X8Y| 5X6Y|5X7Y|5X8Y|
		               |----|----|----| ----|----|----| ----|----|----|
			square7	   |6X0Y|6X1Y|6X2Y| 7X0Y|7X1Y|7X2Y| 8X0Y|8X1Y|8X2Y|
		               |----|----|----| ----|----|----| ----|----|----|
			square8	   |6X3Y|6X4Y|6X5Y| 7X3Y|7X4Y|7X5Y| 8X3Y|8X4Y|8X5Y|
		               |----|----|----| ----|----|----| ----|----|----|
			square9	   |6X6Y|6X7Y|6X8Y| 7X6Y|7X7Y|7X8Y| 8X6Y|8X7Y|8X8Y|
		               |----|----|----| ----|----|----| ----|----|----|
	*/
	xxbase := 0
	for i := 0; i < 3; i++ {
		xbase := xxbase
		ybase := 0
		for i := 0; i < 3; i++ {
			mainarr := GetMainArr()
			x := xbase
			y := ybase
			for i := 0; i < 3; i++ { // y axis for one square
				for i := 0; i < 3; i++ { // x axis for one square
					if mainarr[m[x][y]].exist {
						return false
					}
					mainarr[m[x][y]].exist = true
					y += 1
				}
				x += 1
				y = ybase
			}
			ybase += 3
			y = ybase
		}
		xxbase += 3
		xbase = xxbase
	}

	return true
}
