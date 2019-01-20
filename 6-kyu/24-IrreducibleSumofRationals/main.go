package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	fmt.Println(SumFracts([][]int{{1, 3}, {5, 3}}))
	fmt.Println(SumFracts([][]int{{1, 2}, {1, 3}, {1, 4}}))
	fmt.Println(gcd(36, 24))
	fmt.Println(gcd(81345, 15786))
	fmt.Println(gcd(87546, 11111111))
	fmt.Println(gcd(43216, 255689))
}

// SumFracts :
func SumFracts(arr [][]int) string {
	if len(arr) < 1 {
		return "0"
	}
	res := arr[0]
	for _, v := range arr[1:] {
		res = add(res, v)
	}
	r := float64(res[0]) / float64(res[1])
	if r == math.Trunc(r) {
		return fmt.Sprint(r)
	}
	return fmt.Sprintf("%d/%d", res[0], res[1])
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func add(a []int, b []int) []int {
	s := a[0]*b[1] + a[1]*b[0]
	w := a[1] * b[1]
	d := gcd(s, w)
	return []int{s / d, w / d}
}

/*

 */

// SumFractsB :
func SumFractsB(arr [][]int) string {
	sum := big.NewRat(0, 1)

	for _, f := range arr {
		sum.Add(sum, big.NewRat(int64(f[0]), int64(f[1])))
	}

	if sum.Denom().Cmp(big.NewInt(1)) == 0 {
		return sum.Num().String()
	}
	return sum.String()
}
