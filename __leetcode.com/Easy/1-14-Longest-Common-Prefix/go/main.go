package main

import (
	"fmt"
)

func main() {
	fmt.Println(LongestCommonPrefix([]string{"flower", "flow", "flight"})) // "fl"
}

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	isContains := func(item byte, index int, strs []string) bool {
		for _, word := range strs {
			if index >= len(word) || word[index] != item {
				return false
			}
		}
		return true
	}
	longest := ""
	for i := 0; i < len(strs[0]); i++ {
		if isContains(strs[0][i], i, strs) {
			longest += string(strs[0][i])
		} else {
			break
		}
	}
	return longest
}

// ------------------ LongestCommonPrefix2 ------------------

func LongestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var lastPrefix string = ""
	for v := 0; ; v++ {
		for h := 0; h <= len(strs)-2; h++ {
			if len(strs[h]) < v || len(strs[h+1]) < v || strs[h][:v] != strs[h+1][:v] {
				return lastPrefix
			}
		}
		lastPrefix = strs[0][:v]
	}
	return lastPrefix
}
