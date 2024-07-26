package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for m := i + 1 ; m <= '9' ; m++ {
			for a := m + 1 ; a <= '9' ; a++ {
				z01.PrintRune(i)

			}
		}	
	}
}
