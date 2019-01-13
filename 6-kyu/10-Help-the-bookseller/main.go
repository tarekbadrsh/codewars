package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"A", "B", "C", "D"}

	fmt.Println(StockList(b, c))
	fmt.Println(StockListB(b, c))
}

// StockList :
func StockList(listArt []string, listCat []string) string {
	m := map[string]int{}
	if len(listArt) == 0 {
		return ""
	}

	for _, v := range listArt {
		i := strings.Split(v, " ")[1]
		l, _ := strconv.Atoi(i)
		m[string(v[0])] += l
	}

	var buffer bytes.Buffer
	for k, c := range listCat {
		v := m[c]
		if k != 0 {
			buffer.WriteString(" ")
		}
		buffer.WriteString("(")
		buffer.WriteString(c)
		buffer.WriteString(" : ")
		buffer.WriteString(strconv.Itoa(v))
		buffer.WriteString(")")
		if k != len(listCat)-1 {
			buffer.WriteString(" -")
		}
	}

	return buffer.String()
}

//***B***// other solution

// StockListB :
func StockListB(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}
	counts := make(map[string]int)

	for _, book := range listArt {
		var code string
		var amount int
		if _, err := fmt.Sscanf(book, "%s %d", &code, &amount); err != nil {
			panic("invalid input")
		}
		counts[string(code[0])] += amount
	}

	var collect []string
	for _, cat := range listCat {
		collect = append(collect, fmt.Sprintf("(%s : %d)", cat, counts[cat]))
	}
	return strings.Join(collect, " - ")
}
