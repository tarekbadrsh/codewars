package main

import "fmt"

func main() {
	fmt.Println(comb([]int{4, 3, 5, 6, 10, 20}))
}

func comb(arr []int) int {
	return 0
}

func calc(arr []int) []int {
	for i := 0; i < 6; i++ {

	}
	return arr
}

func getmin(arr []int) (index, val int) {
	m := arr[0]
	i := 0
	for ; i < len(arr); i++ {
		if arr[i] < m {
			m = arr[i]
		}
	}
	return i, m
}

/*
0:[4,3,5,6,10,20]
3:[7,5,6,10,20]
5:[7,11,10,20]
7:[17,11,20]
11:[28,20]
20:[48]
68:[]
*/
