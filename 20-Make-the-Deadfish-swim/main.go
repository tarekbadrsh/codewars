package main

import "fmt"

func main() {
	fmt.Println(Parse("ioioio"))
	fmt.Println(Parse("idoiido"))
	fmt.Println(Parse("isoisoiso"))
	fmt.Println(Parse("codewars"))
}

// Parse :
// i : +1
// d : -1
// s : ^2
// o : return
func Parse(data string) []int {
	res := []int{}
	i := 0
	for _, v := range data {
		switch v {
		case 'o':
			res = append(res, i)
		case 'i':
			i++
		case 'd':
			i--
		case 's':
			i *= i
		}
	}
	return res
}
