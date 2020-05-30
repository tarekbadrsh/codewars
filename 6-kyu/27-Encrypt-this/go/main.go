package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(EncryptThis2("Hello"))
}

// EncryptThis :
func EncryptThis(text string) string {
	words := strings.Split(text, " ")
	result := []string{}
	for _, w := range words {
		nw := ""
		for i, l := range w {
			switch i {
			case 0:
				nw += fmt.Sprint(l)
			case 1:
				nw += string(w[len(w)-1])
			case (len(w) - 1):
				nw += string(w[1])
			default:
				nw += string(l)
			}
		}
		result = append(result, nw)
	}
	return strings.Join(result, " ")
}

// EncryptThis2 :
func EncryptThis2(text string) string {
	wordStart := 0
	wordLength := 0
	var buffer bytes.Buffer
	word := ""
	for i, letter := range text {
		if letter == ' ' || i == len(text)-1 {
			if wordStart == 0 && i == len(text)-1 {
				word = text[wordStart : wordStart+wordLength+1]
			} else if wordStart == 0 || i == len(text)-1 {
				word = text[wordStart : wordStart+wordLength]
			} else {
				word = text[wordStart : wordStart+wordLength-1]
			}
			for i, w := range word {
				if i == 0 {
					buffer.WriteString(fmt.Sprint(w))
				} else if i == 1 {
					buffer.WriteString(string(word[len(word)-1]))
				} else if i == len(word)-1 {
					buffer.WriteString(string(word[1]))
				} else {
					buffer.WriteString(string(w))
				}
			}
			if letter == ' ' {
				buffer.WriteString(" ")
			}
			wordStart = i + 1
			wordLength = 0
		}
		wordLength++
	}
	return buffer.String()
}
