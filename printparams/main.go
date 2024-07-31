package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 0 {
		for _, args := range os.Args {
			argTable := []rune(args)
			for index := range argTable {
				z01.PrintRune(argTable[index])
			}
			z01.PrintRune('\n')
		}

	}
}
