package main

import (
	"github.com/01-edu/z01"
)

const g = "x = 42, y = 21"

func setPoint(ptr *Point) {
	a := 42
	b := 21
	ptr.x = a
	ptr.y = b
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

type Point struct {
	x int
	y int
}

func main() {
	a := 0
	b := 0
	points := &Point{x: a, y: b}
	setPoint(points)
	printStr(g)
}
