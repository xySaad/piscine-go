package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}

	var result int = nb

	var doPower func(int)
	doPower = func(i int) {
		result *= nb
		if i < power-1 {
			doPower(i + 1)
		}
	}

	doPower(1)
	return result
}
