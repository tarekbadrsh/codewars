package main

import (
	"fmt"
)

func main() {
	fmt.Println(DirReduc([]string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}))
	fmt.Println(DirReduc([]string{"NORTH", "SOUTH", "EAST", "WEST"}))
	fmt.Println(DirReduc2([]string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}))
	fmt.Println(DirReduc2([]string{"NORTH", "SOUTH", "EAST", "WEST"}))
	fmt.Println(DirReduc2([]string{"EAST", "SOUTH", "WEST", "EAST", "EAST", "EAST", "EAST", "WEST"}))
}

func DirReduc(arr []string) []string {
	if len(arr) > 1 {
		for i := 0; i < len(arr)-1; i++ {
			if areOpposite(arr[i], arr[i+1]) {
				newarr := append(arr[:i], arr[i+2:]...)
				return DirReduc(newarr)
			}
		}
	}
	return arr
}

func areOpposite(e1, e2 string) bool {
	switch e1 {
	case "NORTH":
		if e2 == "SOUTH" {
			return true
		}
	case "SOUTH":
		if e2 == "NORTH" {
			return true
		}
	case "EAST":
		if e2 == "WEST" {
			return true
		}
	case "WEST":
		if e2 == "EAST" {
			return true
		}
	default:
		return false
	}
	return false
}

func DirReduc2(arr []string) []string {
	for i := len(arr) - 1; i > 0; i-- {
		dA, dB := arr[i], arr[i-1]
		lA, lB := len(arr[i])-1, len(arr[i-1])-1
		if dA[0] != dB[0] && dA[lA] == dB[lB] {
			arr = append(arr[:i-1], arr[i+1:]...)
			i = len(arr)
		}
	}
	return arr
}

func DirReduc3(arr []string) []string {
	zeroMap := map[string]string{
		"NORTH": "SOUTH",
		"SOUTH": "NORTH",
		"WEST":  "EAST",
		"EAST":  "WEST",
	}
	res := make([]string, 0)
	for _, r := range arr {
		l := len(res)
		if l > 0 && res[l-1] == zeroMap[r] {
			res = res[:l-1]
		} else {
			res = append(res, r)
		}
	}
	return res
}
