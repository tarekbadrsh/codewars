package main

import "fmt"

var data []int64

func main() {
	//fmt.Println(Decompose(7))
	calcuateD(50, 50*50)
	fmt.Println(data)
}

// Decompose :
func Decompose(n int64) []int64 {

	return []int64{}
}

func calcuateD(n, r int64) []int64 {
	if r < 1 {
		return []int64{n}
	}

	for i := n - 1; i > 0; i-- {
		if (r - (i * i)) >= 0 {
			// data = append(data, i)
			a := calcuateD(i, (r - (i * i)))
			if a != nil {
				data = append(data, a[0])
			}
		}
	}

	return nil
}
