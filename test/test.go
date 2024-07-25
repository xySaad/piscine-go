package main

import (
	"fmt"
	"piscine"
)

func main() {
	a := 11
	b := 3
	var div int
	var mod int
	piscine.DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
