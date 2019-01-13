package main

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println(BackwardsPrime(2, 100))
	fmt.Println(BackwardsPrime(9900, 10000))
	fmt.Println(BackwardsPrime(1000, 1100))
}

func reverse(n int) int {
	r := 0
	for n > 0 {
		remainder := n % 10
		r *= 10
		r += remainder
		n /= 10
	}
	return r
}

func ismirror(n int) bool {
	s := strconv.Itoa(n)
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	if (n < 3) || (n%2 == 0) {
		return false
	}
	limit := int(math.Sqrt(float64(n)) + 1.0)
	for i := 3; i <= limit; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func checkreverse(n int) (int, bool) {
	r := ""
	mirror := true
	s := strconv.Itoa(n)
	for i := 0; i < len(s); i++ {
		last := s[len(s)-i-1]
		r += string(last)
		if s[i] != last {
			mirror = false
		}
	}
	num, _ := strconv.Atoi(r)
	return num, mirror
}

// BackwardsPrime :
func BackwardsPrime(start int, stop int) []int {
	res := []int{}
	for i := start; i <= stop; i++ {
		rev, mirror := checkreverse(i)
		if !mirror && big.NewInt(int64(i)).ProbablyPrime(20) {
			if big.NewInt(int64(rev)).ProbablyPrime(20) {
				res = append(res, i)
			}
		}
	}
	if len(res) > 0 {
		return res
	}
	return nil
}
