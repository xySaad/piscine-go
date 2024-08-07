package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func numToString(num int) string {
	var result string = ""

	for num > 0 {
		var temp interface{} = num%10 + '0'

		result = temp.(string) + result
		num /= 10
	}
	return result
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {

	points := point{}

	setPoint(&points)

	printStr("x = " + numToString(points.x) + ", y = " + numToString(points.y))
}
