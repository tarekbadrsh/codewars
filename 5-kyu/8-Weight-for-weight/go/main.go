package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(OrderWeight("103 123 4444 99 2000"))                          // "2000 103 123 4444 99"
	fmt.Println(OrderWeight("2000 10003 1234000 44444444 9999 11 11 22 123")) // "11 11 2000 10003 22 123 1234000 44444444 9999"
	fmt.Println(OrderWeight("47 678 184 191 390 750 731 32 233 614 732 682 273 495 47 190 780 890 572 113 488 83 678 587 213 716 370 955 196 183 360 222 191 161"))
	// 113 32 213 222 161 233 360 190 370 191 191 47 47 614 731 83 183 273 390 732 750 184 572 716 780 196 682 890 495 955 488 587 678 678
}

type str struct {
	val    string
	weight int
}

// OrderWeight :
func OrderWeight(strng string) string {
	if strng == "" {
		return ""
	}
	var strmap = []str{}
	var onestr string = ""
	var oneint rune = 0
	for k, v := range strng {
		if v == ' ' && k != 0 {
			strmap = append(strmap, str{onestr, int(oneint)})
			onestr = ""
			oneint = 0
			continue
		}
		onestr += string(v)
		oneint += v - '0'
		if k == len(strng)-1 {
			strmap = append(strmap, str{onestr, int(oneint)})
		}
	}
	sort.Slice(strmap, func(i, j int) bool {
		if strmap[i].weight != strmap[j].weight {
			return strmap[i].weight < strmap[j].weight
		}
		return strmap[i].val < strmap[j].val
	})
	result := ""
	for _, v := range strmap {
		result += v.val + " "
	}
	return result[:len(result)-1]
}

///================ other practices ==================///

// OrderWeightA :
type WeightList []string

func (a WeightList) Len() int      { return len(a) }
func (a WeightList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a WeightList) Less(i, j int) bool {
	if valueForWeight(a[i]) < valueForWeight(a[j]) {
		return true
	}
	if valueForWeight(a[i]) == valueForWeight(a[j]) {
		return a[i] <= a[j]
	}
	return false
}

func valueForWeight(str string) int {
	val := 0
	for _, c := range str {
		val += int(c - '0')
	}
	return val
}

func OrderWeightA(strng string) string {
	weights := WeightList(strings.Split(strng, " "))
	sort.Stable(weights)
	return strings.Join([]string(weights), " ")
}
