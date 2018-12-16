package main

import "fmt"

func main() {
	s := "aaaxbbbbyyhwawiwjjjwwm"
	fmt.Println(PrinterError(s))
}

// PrinterError :
func PrinterError(s string) string {
	l := map[rune]bool{}
	l['a'] = true
	l['b'] = true
	l['c'] = true
	l['d'] = true
	l['e'] = true
	l['f'] = true
	l['g'] = true
	l['h'] = true
	l['i'] = true
	l['j'] = true
	l['k'] = true
	l['l'] = true
	l['m'] = true

	errors := 0
	for _, v := range s {
		if !l[v] {
			errors++
		}
	}
	return fmt.Sprintf("%d/%d", errors, len(s))
}

// PrinterErrorBestPractices :
func PrinterErrorBestPractices(s string) string {
	e := 0
	for _, c := range s {
		if c < 'a' || c > 'm' {
			e++
		}
	}

	return fmt.Sprintf("%d/%d", e, len(s))
}
