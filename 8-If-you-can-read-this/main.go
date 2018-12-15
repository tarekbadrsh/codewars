package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(ToNato("If you can read!"))
}

// ToNato :
func ToNato(words string) string {
	words = strings.ToUpper(words)

	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu"}

	arr := map[rune]string{}

	for _, v := range nato {
		w1 := v[0]
		arr[rune(w1)] = v
	}
	result := ""
	for _, v := range words {
		if v == 32 { // ignore whitespace
			continue
		}
		if arr[v] == "" { // for punctuation
			result += string(v)
		}
		result += arr[v] + " "
	}

	return result[:len(result)-1]

}

// ToNatoBestPractices :
func ToNatoBestPractices(words string) string {
	nato := []string{
		"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot",
		"Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November",
		"Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor",
		"Whiskey", "Xray", "Yankee", "Zulu",
	}
	charToCharlie := map[rune]string{}
	for _, value := range nato {
		charToCharlie[rune(value[0])] = value
	}

	result := ""
	for _, letter := range words {
		if unicode.IsLetter(letter) {
			result += charToCharlie[unicode.ToUpper(letter)] + " "
		} else if unicode.IsPunct(letter) {
			result += string(letter)
		}
	}
	return strings.TrimSpace(result)
}
