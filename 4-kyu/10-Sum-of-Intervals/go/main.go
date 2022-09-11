package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SumOfIntervalsA([][2]int{{100, 500}, {1, 5}, {1, 4}, {7, 10}, {3, 5}, {1, 4}, {8, 12}}))
	fmt.Println(SumOfIntervals([][2]int{{1, 5}, {6, 10}}))
	fmt.Println(SumOfIntervals([][2]int{{1, 5}, {1, 5}}))
	fmt.Println(SumOfIntervals([][2]int{{1, 4}, {7, 10}, {3, 5}}))
	fmt.Println(SumOfIntervals([][2]int{{1, 4}, {3, 5}}))
	fmt.Println(overlap([2]int{3, 4}, [2]int{1, 3}))
	fmt.Println(setlargest([2]int{3, 4}, [2]int{1, 2}))
}

func betweenNum(num, left, right int) bool {
	return num >= left && num <= right
}

func overlap(first, second [2]int) bool {
	if betweenNum(first[0], second[0], second[1]) {
		return true
	}
	if betweenNum(first[1], second[0], second[1]) {
		return true
	}
	if betweenNum(second[0], first[0], first[1]) {
		return true
	}
	if betweenNum(second[1], first[0], first[1]) {
		return true
	}
	return false
}

func setlargest(first, second [2]int) [2]int {
	result := [2]int{}
	if first[0] < second[0] {
		result[0] = first[0]
	} else {
		result[0] = second[0]
	}

	if first[1] > second[1] {
		result[1] = first[1]
	} else {
		result[1] = second[1]
	}
	return result
}

// SumOfIntervals
func SumOfIntervals(intervals [][2]int) int {
	innerIntervals := map[[2]int]struct{}{}
	for i := 0; i < len(intervals); i++ {
		for x := 0; x < len(intervals); x++ {
			if overlap(intervals[i], intervals[x]) {
				intervals[i] = setlargest(intervals[i], intervals[x])
			}
		}
	}
	for i := 0; i < len(intervals); i++ {
		for x := 0; x < len(intervals); x++ {
			if overlap(intervals[i], intervals[x]) {
				intervals[i] = setlargest(intervals[i], intervals[x])
			}
		}
	}
	result := 0
	for _, v := range intervals {
		if _, isExist := innerIntervals[v]; !isExist {
			innerIntervals[v] = struct{}{}
			result += v[1] - v[0]
		}
	}
	return result
}

///================ other practices ==================///
func SumOfIntervalsA(intervals [][2]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	sum := 0
	low := intervals[0][0]
	for _, v := range intervals {
		if v[1] >= low {
			sum += v[1]
			if v[0] > low {
				sum -= v[0]
			} else {
				sum -= low
			}
			low = v[1]
		}
	}

	return sum
}
