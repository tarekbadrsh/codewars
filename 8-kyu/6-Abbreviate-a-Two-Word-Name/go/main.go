package main

func main() {
	AbbrevName("test best")
	AbbrevName("Sam Harris")

}

func toUpper(a byte) byte {
	if a >= 'a' && a <= 'z' {
		return a - ('a' - 'A')
	}
	return a
}

// AbbrevName :
func AbbrevName(name string) string {
	result := []byte{toUpper(name[0]), '.'}
	for i, c := range name {
		if c == ' ' {
			result = append(result, toUpper(name[i+1]))
			break
		}
	}
	return string(result)
}

///================ best practices ==================///

// AbbrevNameB :
func AbbrevNameB(name string) string {
	i := 0

	for name[i] != ' ' {
		i++
	}

	return string(name[0]&95) + "." + string(name[i+1]&95)
}

// AbbrevNameC :
func AbbrevNameC(name string) string {
	i := 0

	for name[i] != ' ' {
		i++
	}

	return string([]byte{name[0] & 95, '.', name[i+1] & 95})
}
