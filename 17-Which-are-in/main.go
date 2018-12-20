package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var a1 = []string{"live", "arp", "strong"}
	var a2 = []string{"lively", "alive", "harp", "sharp", "armstrong"}
	r1 := InArray(a1, a2)
	fmt.Println("Strings:", r1)
}

// InArray :
func InArray(array1 []string, array2 []string) []string {
	res := []string{}
	m := map[string]bool{}

	for _, a := range array1 {
		for _, b := range array2 {
			if strings.Contains(b, a) && !m[a] {
				res = append(res, a)
				m[a] = true
				break
			}
		}
	}
	sort.Strings(res)
	return res
}
