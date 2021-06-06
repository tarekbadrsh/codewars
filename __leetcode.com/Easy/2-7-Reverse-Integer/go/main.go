package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-456))
	fmt.Println(reverse(2147483641))
	fmt.Println(reverse(2147483648))
	fmt.Println(reverse(-2147483648))
}

func reverse(x int) int {
	res := 0
	if x < 0 {
		res, _ = strconv.Atoi(reverseStr(strconv.Itoa(x)[1:]))
		if -res < -2147483648 {
			return 0
		}
		return -res
	}
	res, _ = strconv.Atoi(reverseStr(strconv.Itoa(x)))
	if res > 2147483648 {
		return 0
	}
	return res
}

func reverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
