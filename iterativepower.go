package piscine

func IterativePower(nb int, power int) int {
	if nb < 0 {
		return 0
	}

	var result int = nb

	for i := 1; i < power; i++ {
		result *= nb
	}

	return result
}
