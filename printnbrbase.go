package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		printStr("NV")
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}
	baseLength := len(base)
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	var result []rune
	for nbr > 0 {
		remainder := nbr % baseLength
		result = append(result, rune(base[remainder]))
		nbr /= baseLength
	}

	for i := len(result) - 1; i >= 0; i-- {
		z01.PrintRune(result[i])
	}
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	charMap := make(map[rune]bool)
	for _, char := range base {
		if char == '+' || char == '-' || charMap[char] {
			return false
		}
		charMap[char] = true
	}
	return true
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
