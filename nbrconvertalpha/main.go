package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	var numberAsInteger int = 0
	var sign int = 1
	for i := range s {
		var num int = 0
		if s[i] >= '0' && s[i] <= '9' {
			num = int(s[i] - '0')
		} else if i == 0 {
			switch s[i] {
			case '-':
				sign = -1
			case '+':
				sign = 1
			default:
				return 0
			}
		} else {
			return 0
		}
		numberAsInteger = numberAsInteger*10 + num
	}
	return numberAsInteger * sign
}

func main() {
	args := os.Args[1:]
	isUpper := false

	if os.Args[1] == "--upper" {
		isUpper = true
		args = os.Args[2:]
	}

	for i := range args {
		num := Atoi(args[i])
		if num >= 1 && num <= 25 {
			if isUpper {
				z01.PrintRune(rune('A' + num - 1))
			} else {
				z01.PrintRune(rune('a' + num - 1))
			}
		} else {
			z01.PrintRune(' ')
		}
	}
}
