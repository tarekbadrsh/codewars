package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(Prod2Sum(1, 2, 1, 3))     // 50 // {{1, 7}, {5, 5}}
	fmt.Println(Prod2Sum(2, 3, 4, 5))     // {{2, 23}, {7, 22}}
	fmt.Println(Prod2Sum(1, 2, 2, 3))     // {{1, 8}, {4, 7}}
	fmt.Println(Prod2Sum(1, 1, 3, 5))     // {{2, 8}}
	fmt.Println(Prod2Sum(10, 11, 12, 13)) // 69173 // {{2, 263}, {23, 262}}
	fmt.Println(Prod2Sum(1, 20, -4, -5))  // {{75, 104}, {85, 96}}
	fmt.Println(Prod2Sum(100, 100, 100, 100))
}

// Prod2Sum :
func Prod2Sum(a, b, c, d int) [][]int {
	res := [][]int{}

	n := ((a * a) + (b * b)) * ((c * c) + (d * d))
	l := (n / 2) + 1
	for i := 0; i < l; i++ {
		p := i * i
		if p > n {
			break
		}
		ptmp := math.Sqrt(float64(n - p))
		j := int(ptmp)
		l = j + 1
		if ptmp != math.Trunc(ptmp) {
			continue
		}
		if len(res) > 1 {
			break
		}
		res = append(res, []int{i, j})
	}

	return res
}

// Prod2Sum2 :
//
func Prod2Sum2(a, b, c, d int) [][]int {
	res := make([][]int, 0)
	r1 := []int{abs(a*c + b*d), abs(a*d - b*c)}
	r2 := []int{abs(a*c - b*d), abs(a*d + b*c)}
	sort.Ints(r1)
	sort.Ints(r2)
	res = append(res, r1)
	if r1[0] != r2[0] {
		res = append(res, r2)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
