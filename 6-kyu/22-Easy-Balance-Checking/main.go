package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	book := `
1233.00
125 Hardware;! 24.8?;
123 Flowers 93.5
127 Meat 120.90
120 Picture 34.00
124 Gasoline 11.00
123 Photos;! 71.4?;
122 Picture 93.5
132 Tyres;! 19.00,?;
129 Stamps 13.6
129 Fruits{} 17.6
129 Market;! 128.00?;
121 Gasoline;! 13.6?;
	`
	fmt.Println(Balance(book))
	fmt.Println(BalanceB(book))
	fmt.Println(BalanceC(book))

}

// Balance :
func Balance(book string) string {
	book = strings.TrimSpace(book)
	lines := strings.Split(book, "\n")
	var originalBalance float64
	var expense float64
	var count float64
	arrres := []string{}
	for k, l := range lines {
		for _, v := range l {
			if !(v >= 'A' && v <= 'z') && !(v >= '0' && v <= '9') && v != '.' && v != ' ' {
				l = strings.Replace(l, string(v), "", -1)
			}
		}
		if k == 0 {
			originalBalance, _ = strconv.ParseFloat(l, 10)
			arrres = append(arrres, fmt.Sprintf("Original Balance: %.2f", originalBalance))
		}
		arr := strings.Split(l, " ")
		if len(arr) > 2 {
			count++
			price, _ := strconv.ParseFloat(arr[2], 10)
			expense += price
			originalBalance -= price
			raw := fmt.Sprintf("%v %v %.2f Balance %.2f", arr[0], arr[1], price, originalBalance)
			arrres = append(arrres, raw)
		}
	}
	avg := expense / count
	arrres = append(arrres, fmt.Sprintf("Total expense  %.2f\nAverage expense  %.2f", expense, avg))
	return strings.Join(arrres, "\n")
}

func deleteEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

// BalanceB :
func BalanceB(book string) string {
	re1 := regexp.MustCompile(`[^\n. \dA-Za-z]`)
	bb1 := re1.ReplaceAllString(book, "")
	re2 := regexp.MustCompile(`\n+`)
	bk := deleteEmpty(re2.Split(bb1, -1))
	var total, _ = strconv.ParseFloat(bk[0], 64)
	var current = total
	count := 0
	res := fmt.Sprintf("Original Balance: %.2f", total)
	for _, line := range bk[1:len(bk)] {
		count++
		l := strings.Split(line, " ")
		g, _ := strconv.ParseFloat(l[2], 64)
		current -= g
		res += fmt.Sprintf("\n%s %s %.2f Balance %.2f", l[0], l[1], g, current)
	}
	return res + fmt.Sprintf("\nTotal expense  %.2f\nAverage expense  %.2f", total-current, (total-current)/float64(count))
}

var re1 = regexp.MustCompile(`[^\n. \dA-Za-z]`)
var re2 = regexp.MustCompile(`\n+`)

// BalanceC :
func BalanceC(book string) string {
	book = strings.TrimSpace(book)
	bb1 := re1.ReplaceAllString(book, "")
	lst := re2.Split(bb1, -1)
	res := bytes.Buffer{}
	var expense float64
	var count float64

	originalBalance, _ := strconv.ParseFloat(lst[0], 64)
	fmt.Fprintf(&res, "Original Balance: %.2f\n", originalBalance)

	for _, l := range lst[1:] {
		arr := strings.Split(l, " ")
		if len(arr) > 2 {
			count++
			price, _ := strconv.ParseFloat(arr[2], 64)
			expense += price
			originalBalance -= price
			fmt.Fprintf(&res, "%v %v %.2f Balance %.2f\n", arr[0], arr[1], price, originalBalance)
		}
	}
	avg := expense / count
	fmt.Fprintf(&res, "Total expense  %.2f\nAverage expense  %.2f", expense, avg)
	return res.String()
}
