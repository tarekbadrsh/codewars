package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FromNbToStr(779, []int64{8, 7, 5, 3}))
	fmt.Println(FromNbToStr(187, []int64{8, 7, 5, 3}))
	fmt.Println(FromNbToStr(15, []int64{8, 6, 5}))
	fmt.Println(FromNbToStr(3450, []int64{17, 5, 3}))
}

// Gcd : Recursive function to
// return gcd of a and b
func Gcd(a, b int64) int64 {
	// Everything divides 0
	if a == 0 || b == 0 {
		return 0
	}

	// base case
	if a == b {
		return a
	}

	// a is greater
	if a > b {
		return Gcd(a-b, b)
	}

	return Gcd(a, b-a)
}

// Multiply : Recursive multiply array
func Multiply(arr []int64) int64 {
	if len(arr) == 0 {
		return 1
	}
	if len(arr) == 1 {
		return arr[0]
	}
	return arr[0] * Multiply(arr[1:])
}

// FromNbToStr :
func FromNbToStr(n int64, sys []int64) string {
	if Multiply(sys) < n {
		return "Not applicable"
	}

	res := ""
	for k1, a1 := range sys {
		for _, a2 := range sys[k1+1:] {
			// check for gcd
			if Gcd(a1, a2) != 1 {
				return "Not applicable"
			}
		}
		res += fmt.Sprintf("--%d", n%a1)
	}
	return res[1:] + "-"
}

//***B***//

// GcdB : Recursive function to
// return gcd of a and b
func GcdB(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return GcdB(b, a%b)
}

// Multiplyloop : loop multiply array
func Multiplyloop(arr []int64) int64 {
	var mul int64 = 1
	for i := 0; i < len(arr); i++ {
		mul *= arr[i]
	}
	return mul
}

// FromNbToStrB :
func FromNbToStrB(n int64, sys []int64) string {
	if Multiplyloop(sys) < n {
		return "Not applicable"
	}

	var buffer bytes.Buffer
	buffer.WriteString("-")
	for k1, a1 := range sys {
		for _, a2 := range sys[k1+1:] {
			// check for gcd
			if GcdB(a1, a2) != 1 {
				return "Not applicable"
			}
		}
		if k1 != 0 {
			buffer.WriteString("--")
		}
		buffer.WriteString(strconv.FormatInt(n%a1, 10))
	}
	buffer.WriteString("-")

	return buffer.String()
}
