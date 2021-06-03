package main

import (
	"fmt"
	"sort"
)

// u = [1, 3, 4, 7, 9, 10, 13, 15, 19, 21, 22, 27, ...]
// 1 gives 3 and 4, then 3 gives 7 and 10, 4 gives 9 and 13, then 7 gives 15 and 22 and so on...

func main() {
	fmt.Println(DblLinear(10))   //22
	fmt.Println(DblLinear(50))   //175
	fmt.Println(DblLinear(1000)) //8488
	fmt.Println(DblLinear(100))  //447
}

func SortedInsert(s []int, e int) []int {
	s = append(s, 0)
	L := len(s)
	i := sort.Search(len(s), func(i int) bool { return s[i] > e })
	if i == L { // new value is the biggest
		s[i-1] = e
		return s
	}
	copy(s[i+1:], s[i:])
	s[i] = e
	return s
}

func DblLinear(n int) int { // 3019ms
	result := []int{1}
	base := map[int]bool{}
	for i := 0; ; i++ {
		x1, y1 := (2*result[i] + 1), (3*result[i] + 1)
		if !base[x1] {
			base[x1] = true
			result = SortedInsert(result, x1)
		}
		if !base[y1] {
			base[y1] = true
			result = SortedInsert(result, y1)
		}
		if i >= n {
			break
		}
	}
	return result[n]
}

func DblLinear2(n int) int { // 9280ms
	result := []int{1}
	base := map[int]bool{}
	for i := 0; ; i++ {
		sort.Ints(result)
		x1, y1 := (2*result[i] + 1), (3*result[i] + 1)
		if !base[x1] {
			base[x1] = true
			result = append(result, x1)
		}
		if !base[y1] {
			base[y1] = true
			result = append(result, y1)
		}
		if i >= n {
			break
		}
	}
	return result[n]
}

// --------------- DblLinear3 -----------------
func DblLinear3(n int) int {
	// Code
	u := []int{1}
	i := 0
	j := 0

	var y int
	var z int

	for len(u) <= n {
		y = 2*u[i] + 1
		z = 3*u[j] + 1

		if y < z {
			u = append(u, y)
			i++
		} else if y > z {
			u = append(u, z)
			j++
		} else {
			u = append(u, y)
			i++
			j++
		}
	}
	return u[len(u)-1]
}

// --------------- DblLinear4 -----------------

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func DblLinear4(n int) int {
	array := []int{1}
	y := 0
	z := 0

	for i := 1; i <= n; i++ {
		array = append(array, Min(2*array[y]+1, 3*array[z]+1))
		if array[i] == 2*array[y]+1 {
			y++
		}
		if array[i] == 3*array[z]+1 {
			z++
		}
	}
	return array[n]
}
