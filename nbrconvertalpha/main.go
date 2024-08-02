package main

import (
	"os"
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	isUpper := false

	if os.Args[1] == "--upper" {
		isUpper = true
		args = os.Args[2:]
	}

	for i := range args {
		num := piscine.Atoi(args[i])
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
