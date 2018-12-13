package main

import (
	"fmt"
)

func main() {
	fmt.Println(Going(113))
}

var cache map[float64]float64

func factorial(n float64) float64 {
	if n <= 1 {
		return 1
	}
	if cache[n] == 0 {
		cache[n] = n * factorial(n-1)
	}
	return cache[n]
}

// Going :
// not working starting from 200
func Going(n int) float64 {
	cache = map[float64]float64{}
	input := float64(n)
	var sum float64
	for c := input; c > 0; c-- {
		sum += factorial(c)
	}

	res := (1 / factorial(input)) * sum

	return float64(int64(res*1e6)) / 1e6
}

// GoingBestPractices :
// also not working starting from 200 - with go1.11
func GoingBestPractices(n int) float64 {
	res, inter := 1.0, 1.0
	for i := n; i >= 2; i-- {
		inter = inter * (1.0 / float64(i))
		res += inter
	}
	return float64(int64(res*1e6)) / 1e6
}
