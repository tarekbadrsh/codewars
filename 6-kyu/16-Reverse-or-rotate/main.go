package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(Revrot("563000655734469485", 4))
	fmt.Println(Revrot("1234", 0))
	fmt.Println(Revrot("", 0))
	fmt.Println(Revrot("1234", 5))
	fmt.Println(Revrot("733049910872815764", 5))

}

// Revrot :
func Revrot(s string, n int) string {
	var buffer bytes.Buffer

	if n == 0 || s == "" {
		return buffer.String()
	}
	for n <= len(s) {
		e := s[:n]
		s = s[n:]
		i := 0
		for _, v := range e {
			q, _ := strconv.Atoi(string(v))
			i += q
		}
		if i%2 == 0 {
			for i := len(e) - 1; i >= 0; i-- {
				buffer.WriteString(string(e[i]))
			}
		} else {
			buffer.WriteString(e[1:])
			buffer.WriteString(string(e[0]))
		}
	}
	return buffer.String()
}
