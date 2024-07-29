package piscine

func IsPrime(nb int) bool {

	if nb%2 == 0 {
		return false
	}
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}

	for i := 2; i <= nb; i++ {
		if i != nb {
			if nb%i == 0 {
				return false
			}
		}
	}
	return true
}
