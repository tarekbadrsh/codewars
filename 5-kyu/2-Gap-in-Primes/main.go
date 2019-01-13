package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(Gap(2, 100, 110))
}

// Gap :
func Gap(g, m, n int) []int {
	p := 0
	for i := m; i < n; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(20) {
			if p != 0 && (i-p) == g {
				return []int{p, i}
			}
			p = i
		}
	}
	return nil
}

// Gap2 - work but not performance :
func Gap2(g, m, n int) []int {
	p := 0
	for i := m; i <= n; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				continue
			}
		}

		if isPrime == true {
			if p != 0 && (i-p) == g {
				return []int{p, i}
			}
			p = i
		}
	}
	return nil
}

// Gap3 - work but not performance :
func Gap3(g, m, n int) []int {
	p := 0
	b := map[int]bool{}
	for i := 2; i < n; i++ {
		if b[i] == true {
			continue
		}
		if i > m {
			if p != 0 && (i-p) == g {
				return []int{p, i}
			}
			p = i
		}
		for k := i * i; k < n; k += i {
			b[k] = true
		}
	}
	return nil
}
