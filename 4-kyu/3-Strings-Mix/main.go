package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s1 := "Are the kids at home? aaaaa fffff"
	s2 := "Yes they are here! aaaaa fffff"

	// s1 := "Are they here"
	// s2 := "yes, they are here"

	fmt.Println(Mix(s1, s2))
}

type element struct {
	c      rune
	source string
	count  int
	str    string
}

type elements []element

func (e elements) Len() int {
	return len(e)
}
func (e elements) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e elements) Less(i, j int) bool {
	if e[i].count == e[j].count {
		srci, _ := strconv.Atoi(e[i].source)
		srcj, _ := strconv.Atoi(e[j].source)
		if srci == srcj {
			return e[i].c > e[j].c
		}
		return srci > srcj
	}
	return e[i].count < e[j].count
}

// Mix :
func Mix(s1, s2 string) string {
	ms1 := map[rune]int{}
	for _, v := range s1 {
		if v >= 'a' && v <= 'z' {
			ms1[v]++
		}
	}

	elementsmap := map[rune]element{}
	for k, v := range ms1 {
		if v > 1 {
			elementsmap[k] = element{c: k, source: "1", count: v, str: "1:" + strings.Repeat(string(k), v)}
		}
	}

	ms2 := map[rune]int{}
	for _, v := range s2 {
		if v >= 'a' && v <= 'z' {
			ms2[v]++
		}
	}

	for k, v := range ms2 {
		if v > 1 {
			if va, ok := elementsmap[k]; ok {
				if v > va.count {
					elementsmap[k] = element{c: k, source: "2", count: v, str: "2:" + strings.Repeat(string(k), v)}
				} else if v == va.count {
					elementsmap[k] = element{c: k, source: "3", count: v, str: "=:" + strings.Repeat(string(k), v)}
				}
			} else {
				elementsmap[k] = element{c: k, source: "2", count: v, str: "2:" + strings.Repeat(string(k), v)}
			}
		}
	}

	elementsf := elements{}
	for _, v := range elementsmap {
		elementsf = append(elementsf, v)
	}
	sort.Sort(sort.Reverse(elementsf))
	res := []string{}
	for _, v := range elementsf {
		res = append(res, v.str)
	}
	return strings.Join(res, "/")
}

/***


***/

// MixB :
func MixB(s1, s2 string) string {
	alphabase := "abcdefghijklmnopqrstuvwxyz"
	result := []string{}
	for _, c := range alphabase {
		nbs1 := strings.Count(s1, string(c))
		nbs2 := strings.Count(s2, string(c))
		if nbs1 > 1 || nbs2 > 1 {
			if nbs1 == nbs2 {
				result = append(result, "=:"+strings.Repeat(string(c), nbs1))
			}
			if nbs1 > nbs2 {
				result = append(result, "1:"+strings.Repeat(string(c), nbs1))
			}
			if nbs1 < nbs2 {
				result = append(result, "2:"+strings.Repeat(string(c), nbs2))
			}
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if len(result[i]) == len(result[j]) {
			return result[i] < result[j]
		}
		return len(result[i]) > len(result[j])
	})
	return strings.Join(result, "/")
}
