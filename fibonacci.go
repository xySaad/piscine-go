package piscine

func Fibonacci(index int) int {

	if index < 0 {
		return -1
	}
	num := index - 1
	return (num - 1) + (num - 2)
}
