package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	fmt.Println(Anagrams("abba", []string{"aabb", "abcd", "bbaa", "dada"})) //[]string{"aabb", "bbaa"}
}

func Anagrams(word string, words []string) []string {
	wordmap := map[rune]int{}
	wordlen := len(word)
	for _, v := range word {
		wordmap[v]++
	}
	result := []string{}
	checkwordFunc := func(checkword string) bool {
		if len(checkword) != wordlen {
			return false
		}
		tmpmap := map[rune]int{}
		for _, x := range checkword {
			if wordmap[x] == 0 {
				return false
			}
			tmpmap[x]++
			if tmpmap[x] > wordmap[x] {
				return false
			}
		}
		return true
	}
	for _, checkword := range words {
		angrm := checkwordFunc(checkword)
		if angrm {
			result = append(result, checkword)
		}
	}
	if len(result) < 1 {
		return nil
	}
	return result
}

// ---------------- Anagrams2 ----------------------
func Anagrams2(word string, words []string) []string {
	var result []string
	oringalWordMap := genMap(word)

	for _, word := range words {
		wordMap := genMap(word)
		eq := reflect.DeepEqual(wordMap, oringalWordMap)

		if eq {
			result = append(result, word)
		}
	}
	return result
}

func genMap(word string) map[string]int {
	letters := make(map[string]int)

	for _, letter := range word {
		if _, exists := letters[string(letter)]; exists {
			letters[string(letter)]++

		} else {
			letters[string(letter)] = 1

		}
	}

	return letters
}

// ---------------- Anagrams3 ----------------------
func Anagrams3(word string, words []string) []string {
	r := []rune(word)
	sort.SliceStable(r, func(i, j int) bool { return r[i] < r[j] })
	ref := string(r)
	var o []string
	for _, x := range words {
		h := []rune(x)
		sort.SliceStable(h, func(i, j int) bool { return h[i] < h[j] })
		s := string(h)
		if s == ref {
			o = append(o, x)
		}
	}
	return o
}
