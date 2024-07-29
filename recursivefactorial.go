package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	result := 1

	var factorial func(int)
	factorial = func(index int) {
		if index >= 1 {
			if result > 9223372036854775807/index {
				result = 0
			} else {
				result *= index
				factorial(index - 1)
				index--
			}
		}
	}

	factorial(nb)

	return result
}
