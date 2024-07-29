package piscine

func IsPrime(nb int) bool {
	var result bool = true
	for i := 2; i <= 9; i++ {
		if i != nb {
			if nb%i == 0 {
				result = false
			}
		}
	}
	return result
}
