// https://www.codewars.com/kata/valid-braces/train/go
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(ValidBraces("{}({})[]"))
	fmt.Println(ValidBraces("(({[]{()}}[{(())}[]]))"))
}

// ValidBraces :
func ValidBraces(str string) bool {
	for strings.Index(str, "[]") != -1 || strings.Index(str, "{}") != -1 || strings.Index(str, "()") != -1 {
		str = strings.Replace(str, "[]", "", -1)
		str = strings.Replace(str, "{}", "", -1)
		str = strings.Replace(str, "()", "", -1)
	}
	return len(str) == 0
}

var m = map[string]string{
	"{": "}",
	"(": ")",
	"[": "]",
}

// ValidBracesB :
func ValidBracesB(str string) bool {
	s := make([]string, 0)
	for _, r := range str {
		r := string(r)
		if len(s) > 0 && m[s[len(s)-1]] == r {
			s = s[:len(s)-1]
		} else {
			s = append(s, r)
		}
	}
	return len(s) == 0
}

var re = regexp.MustCompile("\\[\\]|\\(\\)|\\{\\}")

// ValidBracesC :
func ValidBracesC(str string) bool {
	length := len(str)
	for {
		str = re.ReplaceAllLiteralString(str, "")
		if length == len(str) {
			break
		}
		length = len(str)
	}
	return len(str) == 0
}
