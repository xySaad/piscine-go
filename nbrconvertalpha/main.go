package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false

	if len(args) == 0 {
		return
	}

	if args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	for _, arg := range args {
		num := 0
		valid := true

		for _, ch := range arg {
			if ch < '0' || ch > '9' {
				valid = false
				break
			}
			num = num*10 + int(ch-'0')
		}

		if !valid || num < 1 || num > 26 {
			z01.PrintRune(' ')
		} else {
			var letter rune
			if upper {
				letter = 'A' + rune(num-1)
			} else {
				letter = 'a' + rune(num-1)
			}
			z01.PrintRune(letter)
		}
	}
	z01.PrintRune('\n')
}
