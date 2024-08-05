package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune('0' + int32(points.x/10))
	z01.PrintRune('0' + int32(points.x%10))
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	z01.PrintRune('0' + int32(points.y/10))
	z01.PrintRune('0' + int32(points.y%10))
	z01.PrintRune('\n')
}
