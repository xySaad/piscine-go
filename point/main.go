package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func intSize(n int) int {
	size := 1

	for n >= 10 {
		n /= 10
		size *= 10
	}
	return size
}

func intToRunes(n int) (result rune) {
	for i := 0; i < 32; i++ {
		if n&(1<<i) != 0 {
			result |= 1 << i
		}
	}
	return
}

func printStr(table ...interface{}) {

	for _, r := range table {
		switch VType := r.(type) {
		case string:
			for _, v := range VType {
				z01.PrintRune(v)
			}
		case int:
			num := VType
			size := intSize(num)
			for size > 0 {
				z01.PrintRune(intToRunes(num/size%10 + '0'))
				size /= 10
			}
		default:
			printStr("Invalid Type", VType)
		}
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

	printStr("x = ", points.x, ", y = ", points.y)
}
