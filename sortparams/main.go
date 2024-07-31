package main

import (
	"os"

	"github.com/01-edu/z01"
)

func sortRunes(table []string) {
	for i := 0; i < len(table)-1; i++ {
		for i := 0; i < len(table)-1; i++ {
			prev := table[i]
			next := table[i+1]
			if prev > next {
				table[i] = next
				table[i+1] = prev
			}
		}
	}
}

func main() {
	if len(os.Args) > 0 {
		sortRunes(os.Args)
		for argsIndex := 1; argsIndex < len(os.Args); argsIndex++ {
			argTable := []rune(os.Args[argsIndex])
			for index := range argTable {
				z01.PrintRune(argTable[index])
			}
			z01.PrintRune('\n')
		}
	}
}
