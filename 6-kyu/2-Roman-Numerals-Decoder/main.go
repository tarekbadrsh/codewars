package main

import "fmt"

type symbols struct {
}

func main() {
	fmt.Println(Decode(""))
	fmt.Println(Decode("IV"))
	fmt.Println(Decode("XXI"))
	fmt.Println(Decode("MCMXC"))
	fmt.Println(Decode("MMVII"))
	fmt.Println(Decode("MDCLXVI"))
	fmt.Println(Decode("MMMMDIX"))
	fmt.Println(Decode("CLIX"))
	fmt.Println(Decode("MMMXCIII"))
}

// Decode :
func Decode(roman string) int {
	symbols := []struct {
		k string
		v int
	}{
		{k: "M", v: 1000},
		{k: "CM", v: 900},
		{k: "D", v: 500},
		{k: "CD", v: 400},
		{k: "C", v: 100},
		{k: "XC", v: 90},
		{k: "L", v: 50},
		{k: "XL", v: 40},
		{k: "X", v: 10},
		{k: "IX", v: 9},
		{k: "V", v: 5},
		{k: "IV", v: 4},
		{k: "I", v: 1},
	}

	result := 0
	index := 0
	romanlen := len(roman)
	for _, smb := range symbols {
		if romanlen >= index+len(smb.k) {
			for roman[index:index+len(smb.k)] == smb.k {
				result += smb.v
				index += len(smb.k)
				if romanlen < index+len(smb.k) {
					break
				}
			}
		}
	}

	return result
}

var decoder = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// DecodeBestPractices :
func DecodeBestPractices(roman string) int {
	if len(roman) == 0 {
		return 0
	}
	first := decoder[rune(roman[0])]
	if len(roman) == 1 {
		return first
	}
	next := decoder[rune(roman[1])]
	if next > first {
		return (next - first) + DecodeBestPractices(roman[2:])
	}
	return first + DecodeBestPractices(roman[1:])
}

// Decodenotvalid :
func Decodenotvalid(roman string) int {
	symbols := map[rune]int{}
	symbols['I'] = 1
	symbols['V'] = 5
	symbols['X'] = 10
	symbols['L'] = 50
	symbols['C'] = 100
	symbols['D'] = 500
	symbols['M'] = 1000
	result := 0
	if len(roman) == 0 {
		return result
	}
	num0 := symbols[rune(roman[0])]
	result = num0
	if len(roman) == 1 {
		return result
	}
	num1 := symbols[rune(roman[1])]
	i := 1
	if num0 == num1 {
		result += num1
		i = 2
	}
	if len(roman) == 2 {
		if num0 > num1 {
			result = num0 + num1
		} else if num0 < num1 {
			result = num1 - num0
		}
		return result
	}

	for {
		if i == len(roman) {
			break
		}
		n1 := symbols[rune(roman[i])]
		if i+1 == len(roman) {
			result += n1
			break
		}
		n2 := symbols[rune(roman[i+1])]

		if n1 == n2 || n1 > n2 {
			result += (n1 + n2)
		} else if n1 < n2 {
			result += (n2 - n1)
		}
		i += 2
	}

	return result
}
