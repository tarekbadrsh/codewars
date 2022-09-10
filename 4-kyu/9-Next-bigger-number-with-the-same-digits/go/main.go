package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(NextBiggerA(144))
	fmt.Println(NextBigger(414))
	fmt.Println(NextBigger(12))
	fmt.Println(NextBigger(2017))
	fmt.Println(NextBigger(531))
	fmt.Println(NextBigger(111))
	fmt.Println(NextBigger(11112))
	fmt.Println(NextBigger(59884848459853))
}

// NextBigger :
func NextBigger(n int) int {
	if n < 12 {
		return -1
	}
	strbase := strconv.Itoa(n)
	// can't be rearranged
	canBe := false
	for i := 0; i < len(strbase)-1; i++ {
		if strbase[i] < strbase[i+1] {
			canBe = true
			break
		}
	}
	if !canBe {
		return -1
	}
	// can't be rearranged
	for nPlus := n + 1; ; nPlus++ {
		strPlus := strconv.Itoa(nPlus)
		valid := true
		for _, v := range strbase {
			if strings.Count(strbase, string(v)) != strings.Count(strPlus, string(v)) {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}
		return nPlus
	}
}

///================ other practices ==================///

func NextBiggerA(n int) (r int) {
	s := []rune(strconv.Itoa(n))

	i := len(s) - 1
	for i > 0 && s[i-1] >= s[i] {
		i--
	}
	if i == 0 {
		return -1
	}

	j := len(s) - 1
	for s[j] <= s[i-1] {
		j--
	}
	s[i-1], s[j] = s[j], s[i-1]

	for k, m := i, len(s)-1; k < m; k, m = k+1, m-1 {
		s[k], s[m] = s[m], s[k]
	}
	r, _ = strconv.Atoi(string(s))
	return
}
