package main

import "fmt"

func main() {
	fmt.Println(SumEvenFibonacci(111111))
}

//
// solution 1
//
// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first := 0
	second := 1
	var myFn = func() int {
		result := first
		temp := first
		first = second
		second = temp + second

		return result
	}
	return myFn
}

// SumEvenFibonacci :
func SumEvenFibonacci(limit int) int {
	if limit == 1 {
		return 2
	}
	res := 0
	f := fibonacci()
	for i := 0; true; i++ {
		a := f()
		if a > limit {
			break
		}
		if a%2 == 0 {
			res += a
		}
	}

	return res
}

// SumEvenFibonacciBestPractices :
func SumEvenFibonacciBestPractices(limit int) int {
	sum, a, b := 2, 1, 2
	for b <= limit {
		a, b = b, a+b
		if b%2 == 0 {
			sum += b
		}
	}
	return sum
}
