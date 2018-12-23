package main

import "fmt"

func main() {
	fmt.Println(SquaresInRect(3, 5))
	fmt.Println(SquaresInRect(20, 14))

}

// SquaresInRect :
func SquaresInRect(lng int, wdth int) []int {
	if lng == wdth {
		return nil
	}
	res := []int{}
	for {
		if lng <= 0 || wdth <= 0 {
			break
		}
		if lng < wdth {
			res = append(res, lng)
			wdth -= lng
		} else {
			res = append(res, wdth)
			lng -= wdth
		}
	}
	return res
}
