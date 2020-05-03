package main

import "math"

func main() {
	MapVector([]float64{5, 5}, []float64{10, 20, 42}, []float64{-100, -42, 69})
}

// MapVector :
func MapVector(vector, circle1, circle2 []float64) []float64 {
	// r = r1/r2
	r := circle2[2] / circle1[2]
	// xp = x12 - x11
	xp1 := circle1[0] - vector[0]
	// yp = y12 - y11
	yp1 := circle1[1] - vector[1]

	// x2 = x22 - x21
	x2 := circle2[0] - (r * xp1)
	x2 = math.Round(x2*100) / 100

	// y2 = y22 - y21
	y2 := circle2[1] - (r * yp1)
	y2 = math.Round(y2*100) / 100

	return []float64{x2, y2}
}

///================ other practices ==================///

// MapVectorB :
func MapVectorB(vector, circle1, circle2 []float64) []float64 {
	return []float64{circle2[0] + ((vector[0] - circle1[0]) * circle2[2] / circle1[2]), circle2[1] + ((vector[1] - circle1[1]) * circle2[2] / circle1[2])}
}
