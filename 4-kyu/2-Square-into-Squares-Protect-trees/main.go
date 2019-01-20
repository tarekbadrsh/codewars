package main

import "fmt"

func main() {
	fmt.Println(Decompose(7))
}

// Decompose :
func Decompose(n int64) []int64 {
	r := calcuateD(n, n*n)
	if len(r) > 0 {
		return r[:len(r)-1]
	}
	return []int64{}
}

func calcuateD(n, r int64) []int64 {
	if r < 1 {
		return []int64{n}
	}

	for i := n - 1; i > 0; i-- {
		if (r - (i * i)) >= 0 {
			a := calcuateD(i, (r - (i * i)))
			if a != nil {
				a = append(a, n)
				return a
			}
		}
	}

	return nil
}

/*


 */

// DecomposeB :
func DecomposeB(n int64) []int64 {
	return loop(n*n, n)
}

func loop(s int64, i int64) []int64 {
	if s < 0 {
		return nil
	}

	if s == 0 {
		return []int64{}
	}

	for j := i - 1; j > 0; j-- {
		var sub = loop(s-j*j, j)
		if sub != nil {
			return append(sub, []int64{j}...)
		}
	}
	return nil
}
