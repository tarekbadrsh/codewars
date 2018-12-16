package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(FindShort("TqIvISSWVgJX DgFGpLqBRcM UFC WNJQwePdHZDlqJpSmkzDMWukHPYqANYGCRYBLz YYMOodYmn uqllvYtnfSZTAGZQhLNvJM EazmCLnjgkQjVKIatVfxazCQv spvgknomkrDvaPGLCRouOtWWxJCFqHMkFYIFcanNdKwCDCgWaNCSrEes mmiiNHLWnmBfbnltSnvmfgsOizLusoOogMaPsSnUXDOii HFAidV CVZfwaLVSXaAVWeJjslOXQHxSYDCqTLFhfeNrzDcebjUVlTdkffxwMqRxvEGB aKNbDziuBSIAVRGmkBhEUMfkWtPgeJDXWNMNrrlnTZqzYJkWL"))
}

// FindShort :
// solve the problem without using split
// much more performance effective
func FindShort(s string) int {
	res := len(s)
	tmp := 0

	for k, v := range s {
		last := (len(s) == k+1)
		if last {
			tmp++
		}
		if v == ' ' || last {
			if tmp < res {
				res = tmp
			}
			tmp = -1
		}
		tmp++
	}

	return res
}

// FindShort2 :
func FindShort2(s string) int {
	arr := strings.Split(s, " ")
	min := len(s)
	for _, v := range arr {
		if len(v) < min {
			min = len(v)
		}
	}
	return min
}
