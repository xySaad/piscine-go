package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 0 {
		for i := 2; i < len(os.Args[0]); i++ {
			z01.PrintRune(rune(os.Args[0][i]))
		}
		z01.PrintRune('\n')
	}
}
