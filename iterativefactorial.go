package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	result := 1

	for i := nb; i >= 1; i-- {
		if result > 147483647/i {
			return 0

		}
		result *= i
	}

	return result
}
