package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	stringLength := 0

	for range s {
		stringLength++
	}

	for i := 0; i < stringLength; i++ {
		z01.PrintRune(rune(s[i]))
	}
}
