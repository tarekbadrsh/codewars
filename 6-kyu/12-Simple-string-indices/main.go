package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(Solution("((1)23(45))(aB)", 0))
	fmt.Println(Solution("((1)23(45))(aB)", 1))
	fmt.Println(Solution("((1)23(45))(aB)", 6))
	fmt.Println(Solution("((>)|?(*'))(yZ)", 11))
	fmt.Println(Solution(`(|"M=-H;!(.(l)Mea)(&f<+)C(q)$)OTv8K`, 23))
	fmt.Println(Solution("(W&C()1)o()().", 10))

}

// Solution :
func Solution(str string, idx uint) (uint, error) {
	i := 0
	for k, v := range str[idx:] {
		switch v {
		case '(':
			i++
		case ')':
			i--
		}
		if i <= 0 {
			if k == 0 {
				break
			}
			if idx != 0 {
				return uint(k) + idx, nil
			}
			return uint(k), nil
		}
	}

	return 0, errors.New("Not an opening bracket")
}
