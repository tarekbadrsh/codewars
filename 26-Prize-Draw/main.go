package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	st := "Elijah,Chloe,Elizabeth,Matthew,Natalie,Jayden"
	we := []int{1, 3, 5, 5, 3, 6}
	n := 2
	fmt.Println(NthRank(st, we, n))

}

type participate struct {
	name string
	w    int
	corr int
	rank int
}

type arrparticipate []participate

func (p arrparticipate) Len() int {
	return len(p)
}
func (p arrparticipate) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p arrparticipate) Less(i, j int) bool {
	if p[i].rank == p[j].rank {
		// a := p[i].name
		// b := p[j].name
		// for ai := 0; ai < len(a); ai++ {
		// 	if len(b) < ai {
		// 		return false
		// 	}
		// 	if a[ai] == b[ai] {
		// 		continue
		// 	}
		// 	return a[ai] > b[ai]
		// }
		return p[i].name > p[j].name
	}
	return p[i].rank < p[j].rank
}

var english = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26}

// NthRank :
func NthRank(st string, we []int, n int) string {
	if st == "" {
		return "No participants"
	}
	arr := strings.Split(st, ",")

	particis := arrparticipate{}

	for k, v := range arr {
		p := participate{
			name: v,
			w:    we[k],
		}
		p.corr += len(p.name)
		for _, l := range v {
			p.corr += english[strings.ToLower(string(l))]
		}
		p.rank = p.w * p.corr
		particis = append(particis, p)
	}
	sort.Sort(sort.Reverse(particis))

	if len(particis) >= n {
		return particis[n-1].name
	}
	return "Not enough participants"
}

//***solution B***//

type rankStruct struct {
	name string
	wght int
}

// NthRankB :
func NthRankB(st string, we []int, n int) string {
	if st == "" {
		return "No participants"
	}
	names := strings.Split(st, ",")
	if n > len(names) {
		return "Not enough participants"
	}
	ranks := make([]rankStruct, len(names))
	for i := range ranks {
		ranks[i].name = names[i]
		w := len(names[i])
		for _, c := range strings.ToLower(names[i]) {
			w += int(c) - 'a' + 1
		}
		w *= we[i]
		ranks[i].wght = w
	}
	sort.Slice(ranks, func(i, j int) bool {
		if ranks[i].wght == ranks[j].wght {
			return ranks[i].name < ranks[j].name
		}
		return ranks[i].wght > ranks[j].wght
	})
	return ranks[n-1].name
}
