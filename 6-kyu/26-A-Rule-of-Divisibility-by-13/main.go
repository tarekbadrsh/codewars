package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1, 10, 9, 12, 3, 4
	// 5634
	// 4365

	// a := (4 * 1) + (3 * 10) + (6 * 9) + (5 * 12) //148
	// fmt.Println(a)
	// b := (8 * 1) + (4 * 10) + (1 * 9) //57
	// fmt.Println(b)
	// c := (7 * 1) + (5 * 10) //57
	// fmt.Println(c)
	fmt.Println(Thirt(5634))     // 57
	fmt.Println(Thirt(85299258)) // 31
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Thirt :
func Thirt(n int) int {
	arr := []int{1, 10, 9, 12, 3, 4}
	sum, res := 0, 0
	s := fmt.Sprint(n)
	for {
		s = reverse(s)
		for k, v := range s {
			a, _ := strconv.Atoi(string(v))
			res += arr[k%len(arr)] * a
		}
		if sum == res {
			return res
		}
		s = fmt.Sprint(res)
		sum = res
		res = 0
	}
}

// ThirtB :
func ThirtB(n int) int {
	w := [6]int{1, 10, 9, 12, 3, 4}
	for {
		r, q, c, j := n, -1, 0, 0
		for q != 0 {
			q = r / 10
			c += (r % 10) * w[j%6]
			r = q
			j++
		}
		if c == n {
			return c
		}
		n = c
	}
}
