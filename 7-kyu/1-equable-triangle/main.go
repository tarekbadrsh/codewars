package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(EquableTriangle(2, 3, 4))
}

// EquableTriangle :
func EquableTriangle(a, b, c int) bool {
	// perimeter is a + b + c
	perimeter := a + b + c
	// where p is half the perimeter
	p := perimeter / 2
	// Area = Sqrt(p * ((p - a) * (p - b) * (p - c)))
	area := math.Sqrt(float64(p * ((p - a) * (p - b) * (p - c))))
	return (float64(perimeter) == area)
}

// EquableTriangleBestPractices :
func EquableTriangleBestPractices(a, b, c int) bool {
	p := a + b + c
	return 16*p == (p-2*a)*(p-2*b)*(p-2*c)
}
