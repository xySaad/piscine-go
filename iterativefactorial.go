package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	result := 1

	for i := nb; i >= 1; i-- {
		if result > 9223372036854775807/i {
			return 0
		}
		result *= i
	}

	return result
}
