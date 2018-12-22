package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	l1 := "GOOG 300 542.0 B, AAPL 50 145.0 B, CSCO 250.0 29 B, GOOG 200 580.0 S"
	fmt.Println(BalanceStatement(l1))
	l2 := "ZNGA 1300 2.66 B, CLH15.NYM 50 56.32 B, OWW 1000 11.623 B, OGG 20 580.1 B"
	fmt.Println(BalanceStatement(l2))
	l3 := "CLH15.NYM 10 5 B"
	fmt.Println(BalanceStatement(l3))
}

// BalanceStatement :
func BalanceStatement(lst string) string {
	if len(lst) == 0 {
		return "Buy: 0 Sell: 0"
	}
	var badform = []string{}
	prices := map[string]float64{"B": 0.0, "S": 0.0}
	arr := strings.Split(lst, ", ")
	var valid = regexp.MustCompile(`\S+ \d+ \d*\.\d+ (B|S)`)
	for _, v := range arr {
		if !valid.MatchString(v) {
			badform = append(badform, v+" ;")
		} else {
			uu := strings.Split(v, " ")
			_, quantity, price, status := uu[0], uu[1], uu[2], uu[3]
			q, _ := strconv.Atoi(quantity)
			p, _ := strconv.ParseFloat(price, 64)
			newprice := prices[status] + float64(q)*p
			prices[status] = newprice
		}
	}
	res := fmt.Sprintf("Buy: %.0f Sell: %.0f", prices["B"], prices["S"])
	if len(badform) != 0 {
		res += fmt.Sprintf("; Badly formed %d: %s", len(badform), strings.Join(badform, ""))
	}
	return res
}

/*
// another solution

func isBadlyFormed(line string) bool {
	fields := strings.Split(line, " ")
	if len(fields) == 4 {
		if fields[3] == "B" || fields[3] == "S" {
			if strings.Contains(fields[2], ".") {
				return false
			}
		}
	}
	return true
}

// BalanceStatement :
func BalanceStatement(lst string) string {
	if lst == "" {
		return "Buy: 0 Sell: 0"
	}
	var buy float64
	var sell float64
	var badlyformed string
	var badly int
	arr := strings.Split(lst, ", ")
	for _, line := range arr {
		if isBadlyFormed(line) {
			badlyformed += line + " ;"
			badly++
			continue
		}
		fields := strings.Split(line, " ")
		quantity, _ := strconv.ParseFloat(fields[1], 10)
		price, _ := strconv.ParseFloat(fields[2], 10)

		switch fields[3] {
		case "S":
			sell += price * quantity
		case "B":
			buy += price * quantity
		}
	}

	if badly == 0 {
		return fmt.Sprintf("Buy: %v Sell: %v", math.Round(buy), math.Round(sell))
	}
	return fmt.Sprintf("Buy: %v Sell: %v; Badly formed %d: %v", math.Round(buy), math.Round(sell), badly, badlyformed)
}
*/

/*
// another solution

// BalanceStatement :
func BalanceStatement(lst string) string {
	var buy float64
	var sell float64
	badly := 0
	badlyformed := ""
	arr := strings.Split(lst, ", ")
	if len(arr) == 1 && strings.TrimSpace(arr[0]) == "" {
		return "Buy: 0 Sell: 0"
	}
	for _, order := range arr {
		var code string
		var quantity int
		var price string
		var status string
		_, err := fmt.Sscanf(order, "%s %d %s %s", &code, &quantity, &price, &status)
		if err != nil {
			badly++
			badlyformed += fmt.Sprintf("%v ;", order)
			continue
		}
		if !strings.Contains(price, ".") {
			badly++
			badlyformed += fmt.Sprintf("%v ;", order)
			continue
		}
		switch status {
		case "S":
			v, _ := strconv.ParseFloat(price, 10)
			sell += v * float64(quantity)
		case "B":
			v, _ := strconv.ParseFloat(price, 10)
			buy += v * float64(quantity)
		default:
			badly++
			badlyformed += fmt.Sprintf("%v ;", order)
			continue
		}
	}
	badlyformed = fmt.Sprintf("Badly formed %d: %v", badly, badlyformed)

	if badly == 0 {
		return fmt.Sprintf("Buy: %v Sell: %v", math.Round(buy), math.Round(sell))
	}
	return fmt.Sprintf("Buy: %v Sell: %v; %v", math.Round(buy), math.Round(sell), badlyformed)
}*/
