package kata

import (
	"strconv"
	"strings"
)

// DecipherThis :
func DecipherThis(text string) string {
	words := strings.Split(text, " ")
	result := []string{}
	for _, w := range words {
		nw := ""
		ll := ""
		lenOfNumber := 0
		for i, letter := range w {
			if letter >= 48 && letter <= 57 {
				ll += string(letter)
				lenOfNumber++
				if lenOfNumber == len(w) {
					nn, _ := strconv.Atoi(ll)
					nw += string(nn)
					ll = ""
				}
			} else {
				if ll != "" {
					nn, _ := strconv.Atoi(ll)
					nw += string(nn)
					ll = ""
				}
				if i == lenOfNumber {
					nw += string(w[len(w)-1])
				} else if i == len(w)-1 {
					nw += string(w[lenOfNumber])
				} else {
					nw += string(w[i])
				}
			}
		}
		result = append(result, nw)
	}
	return strings.Join(result, " ")
}
