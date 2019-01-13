package main

import "fmt"

func main() {
	fmt.Println(NbYear(1000, 0.02, 50, 1200))
}

// NbYear :
func NbYear(p0 int, percent float64, aug int, p int) int {
	percent = percent / 100
	c := 0
	_, c = growthPopulation(p0, percent, aug, p, c)
	return c
}

func growthPopulation(p0 int, percent float64, aug int, p int, c int) (int, int) {
	c++
	cy := p0 + int(float64(p0)*percent) + aug
	if cy < p {
		return growthPopulation(cy, percent, aug, p, c)
	}
	return cy, c
}

// NbYearBestPractices :
func NbYearBestPractices(p0 int, percent float64, aug int, p int) int {
	var y int
	for y = 0; p0 < p; y++ {
		p0 = int(float64(p0)*(1.0+percent*0.01)) + aug
	}
	return y
}
