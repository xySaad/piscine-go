package piscine

func FindNextPrime(nb int) int {
	for i := nb; i <= 9223372036854775807; i++ {
		if IsPrime(i) {
			return i
		}
	}
	return 0
}
