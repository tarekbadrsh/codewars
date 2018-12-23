package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	s1 := "6900690040, 4690606946, 9990494604"
	fmt.Println(PosAverage(s1))
	s2 := "466960, 069060, 494940, 060069, 060090, 640009, 496464, 606900, 004000, 944096"
	fmt.Println(PosAverage(s2))
	s3 := "444996, 699990, 666690, 096904, 600644, 640646, 606469, 409694, 666094, 606490"
	fmt.Println(PosAverage(s3))
}

func getcommon(a1, a2 string) int {
	count := 0
	for i := 0; i < len(a1); i++ {
		if a1[i] == a2[i] {
			count++
		}
	}
	return count
}

// PosAverage :
func PosAverage(s string) float64 {
	arr := strings.Split(s, ", ")

	// get combinations count
	arrln := len(arr)
	elemln := len(arr[0])
	n := (arrln * (arrln - 1)) / 2
	combinations := elemln * n

	// get common count
	countcommon := 0
	for k1, v1 := range arr {
		for _, v2 := range arr[k1+1:] {
			countcommon += getcommon(v1, v2)
		}
	}

	res := (float64(countcommon) / float64(combinations)) * 100
	return res
}

//***solution B***//

func pair(str1 string, str2 string) float64 {
	cnt := 0
	for i := range str1 {
		if str1[i] == str2[i] {
			cnt++
		}
	}
	return float64(cnt) / float64(len(str1))
}

// PosAverageB :
func PosAverageB(s string) float64 {
	arr := strings.Split(s, ", ")
	n := 0.0
	cnt := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			n += pair(arr[i], arr[j])
			cnt++
		}
	}
	res := n * 100.0 / float64(cnt)
	return float64(math.Round(res*1e10)) / 1e10
}
