package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(LengthOfLongestSubstring("bbbb"))     //b : 1
	fmt.Println(LengthOfLongestSubstring("abcabcbb")) //abc : 3
	fmt.Println(LengthOfLongestSubstring("pwwkew"))   //wke : 3
	fmt.Println(LengthOfLongestSubstring("aab"))      //ab : 2
	fmt.Println(LengthOfLongestSubstring("dvdfd"))    //vdf : 3
	fmt.Println(LengthOfLongestSubstring("dwavdf"))   //wavd : 5
}

func LengthOfLongestSubstring(s string) int {
	result := ""
	tmpResult := ""
	for i := 0; i < len(s); i++ {
		subs := string(s[i])
		if strings.Contains(tmpResult, subs) {
			a := strings.Index(s, subs)
			s = s[a+1:]
			i = -1
			tmpResult = ""
		} else {
			tmpResult += subs
			if len(tmpResult) >= len(result) {
				result = tmpResult
			}
		}
	}
	return len(result)
}

func LengthOfLongestSubstring2(s string) int {
	loc := [256]int{}
	for i := range loc {
		loc[i] = -1
	}
	length, left := 0, 0
	for i := 0; i < len(s); i++ {
		if loc[s[i]] >= left {
			left = loc[s[i]] + 1
		} else if i+1-left > length {
			length = i + 1 - left
		}
		loc[s[i]] = i
	}
	return length
}

func LengthOfLongestSubstring3(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	m := make(map[rune]int)
	left, result := -1, 0
	for i, v := range s {
		if lastSeen, ok := m[v]; ok && lastSeen > left {
			left = lastSeen
		}
		if result < i-left {
			result = i - left
		}
		m[v] = i
	}
	return result
}
