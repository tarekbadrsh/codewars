package main

import "fmt"

func main() {
	fmt.Println(ValidBraces("{}({})[]"))
	//	fmt.Println(endBraces("(({[]{()}}[{(())}[]]))"))
}

// ValidBraces :
func ValidBraces(str string) bool {
	l := len(str)

	if l%2 != 0 {
		return false
	}

	for i := 0; i < l; i += 2 {
		m, _ := checkBraces(str[i], str[l-i-1])
		if i >= l/2 {
			m = true
		}

		c, _ := checkBraces(str[i], str[i+1])
		if !m && !c {
			return false
		}
	}

	return true
}

// func validateRangeBraces(str string) bool {

// 	return true
// }

// func endBraces(str string) int {
// 	for i := 0; i < l; i += 2 {
// 		c, _ := checkBraces(str[i], str[i+1])
// 		if !m && !c {
// 			return false
// 		}
// 	}
// 	return 0
// }

func checkBraces(a, b byte) (bool, byte) {
	switch a {
	case '{':
		if b != '}' {
			return false, b
		}
		break
	case '[':
		if b != ']' {
			return false, b
		}
		break
	case '(':
		if b != ')' {
			return false, b
		}
		break
	default:
		return false, 0
	}
	return true, 0
}
