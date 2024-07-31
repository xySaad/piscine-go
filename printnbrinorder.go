package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	digits := []int{}
	if n == 0 {
		z01.PrintRune('0')
	}
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	SortIntegerTable(digits)
	for i := range digits {
		z01.PrintRune(rune(digits[i] + '0'))
		n *= 10
		n += digits[i]
	}
}
