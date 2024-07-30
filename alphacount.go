package piscine

func AlphaCount(s string) int {
	count := 0

	for i := 0; i <= len(s); i++ {
		if s[i] <= 'Z' && s[i] >= 'a' {
			count++
		}

	}
	return count
}
