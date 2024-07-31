package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 0 {
		for i := 0; i < len(os.Args[0]); i++ {
			if os.Args[0][i] != '/' && os.Args[0][i] != '.' {
				z01.PrintRune(rune(os.Args[0][i]))
			}
		}
		z01.PrintRune('\n')
	}
}
