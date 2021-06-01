package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(SpinWords("Hey fellow warriors")) //Hey wollef sroirraw
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func SpinWords(str string) string {
	strlist := strings.Split(str, " ")
	result := []string{}
	for _, v := range strlist {
		if len(v) > 4 {
			result = append(result, reverse(v))
			continue
		}
		result = append(result, v)
	}
	return strings.Join(result, " ")
}

//--------- SpinWord2 ------------------
func SpinWords2(str string) string {
	strSplit := strings.Split(str, " ")
	for index, value := range strSplit {
		if len(value) > 4 {
			strSplit[index] = ""
			for _, s := range value {
				strSplit[index] = string(s) + strSplit[index]
			}
		}
	}

	return strings.Join(strSplit, " ")
}
