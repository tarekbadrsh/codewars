package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsValidWalk([]rune{'n', 's'}))
	fmt.Println(IsValidWalk([]rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}))
	fmt.Println(IsValidWalk([]rune{'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e'}))
}

// IsValidWalk :
func IsValidWalk(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}
	counter := map[rune]int{}
	for i := 0; i < len(walk); i++ {
		counter[walk[i]] += 1
	}
	return (counter['n']-counter['s'] == 0) && (counter['w']-counter['e'] == 0)
}

///================ other practices ==================///

func IsValidWalkA(walk []rune) bool {
	if len(walk) != 10 {
		return false
	}
	var x, y int
	for _, r := range walk {
		switch r {
		case 's':
			y--
		case 'n':
			y++
		case 'w':
			x--
		case 'e':
			x++
		}
	}
	return x == 0 && y == 0
}
