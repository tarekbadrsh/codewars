package main

import "fmt"

func main() {
	fmt.Println(GetGrade(95, 90, 93))
	fmt.Println(GetGrade(70, 70, 100))
}

// GetGrade :
func GetGrade(a, b, c int) rune {
	total := a + b + c

	score := total / 3

	if 90 <= score && score <= 100 {
		return 'A'
	}
	if 80 <= score && score <= 90 {
		return 'B'
	}
	if 70 <= score && score <= 80 {
		return 'C'
	}
	if 60 <= score && score <= 70 {
		return 'D'
	}
	if 0 <= score && score <= 60 {
		return 'F'
	}
	return 'A'
}

///================ best practices ==================///

// GetGradeA :
func GetGradeA(a, b, c int) rune {
	switch (a + b + c) / 30 {
	case 10:
		return 'A'
	case 9:
		return 'A'
	case 8:
		return 'B'
	case 7:
		return 'C'
	case 6:
		return 'D'
	default:
		return 'F'
	}
}

// GetGradeB :
func GetGradeB(a, b, c int) rune {
	return rune("FFFFFFDCBAA"[(a+b+c)/30])
}

// GetGradeC :
func GetGradeC(a, b, c int) rune {
	score := map[int]rune{10: 'A', 9: 'A', 8: 'B', 7: 'C', 6: 'D'}
	if s, ok := score[(a+b+c)/30]; ok {
		return s
	}
	return 'F'
}

// GetGradeD :
func GetGradeD(a, b, c int) rune {
	score := map[int]rune{10: 'A', 9: 'A', 8: 'B', 7: 'C', 6: 'D', 5: 'F', 4: 'F', 3: 'F', 2: 'F', 1: 'F', 0: 'F'}
	return score[(a+b+c)/30]
}
