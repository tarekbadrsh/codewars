package main

import "fmt"

func main() {
	fmt.Println(ValidBraces("{}({})[]"))
}

// ValidBraces :
func ValidBraces(str string) bool {
	l := len(str)
	if l%2 != 0 {
		return false
	}
	for i := 0; i < l; i += 2 {
		m := checkBraces(str[i], str[l-i-1])
		if i >= l/2 {
			m = true
		}

		c := checkBraces(str[i], str[i+1])
		if !m && !c {
			return false
		}
	}

	return true
}

func checkBraces(a, b byte) bool {
	switch a {
	case '{':
		if b != '}' {
			return false
		}
		break
	case '[':
		if b != ']' {
			return false
		}
		break
	case '(':
		if b != ')' {
			return false
		}
		break
	default:
		return false
	}
	return true
}
