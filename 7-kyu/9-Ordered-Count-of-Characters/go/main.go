package main

import (
	"fmt"
)

// Tuple :
type Tuple struct {
	Char  rune
	Count int
}

func main() {
	fmt.Println(OrderedCountA("abracadabra"))
}

type tmpTuple struct {
	index       int
	count       int
	resutlIndex int
}

// OrderedCount :
func OrderedCount(text string) []Tuple {
	mres := map[rune]tmpTuple{}
	res := make([]Tuple, len(text))
	newindex := 0
	for index, c := range text {
		v, ok := mres[c]
		if !ok {
			v = tmpTuple{index: index, count: 0, resutlIndex: newindex}
			newindex++
		}
		v.count++
		mres[c] = v

		res[v.resutlIndex] = Tuple{c, v.count}
	}
	return res[:len(mres)]
}

///================ other practices ==================///

// OrderedCountA :
func OrderedCountA(text string) []Tuple {
	result := make([]Tuple, 0, 27)
	indexer := make([]int, 127)
	counter := 1
	for _, c := range text {
		if indexer[c] == 0 {
			indexer[c] = counter
			counter++
		}
		if len(result) <= indexer[c] {
			result = append(result, Tuple{c, 0})
		}
		result[indexer[c]-1].Count++
	}
	return result[:counter-1]
}

// OrderedCountB :
func OrderedCountB(text string) []Tuple {
	dic := make([]int, 127)
	order := make([]rune, 0, 127)
	ret := make([]Tuple, 0, 20)

	for _, c := range text {
		if dic[c] == 0 {
			order = append(order, c)
		}
		dic[c]++
	}
	for _, c := range order {
		ret = append(ret, Tuple{c, dic[c]})
	}
	return ret
}
