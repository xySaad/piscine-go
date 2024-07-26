package piscine

func BasicAtoi(s string) int {
	var numberAsInteger int = 0

	for i := range s {
		alo := s[i] - '0'
		numberAsInteger = numberAsInteger*10 + alo
	}

	return numberAsInteger
}
