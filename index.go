package piscine

func Index(s string, toFind string) int {
	if len(toFind) == 0 {
		return 0
	}

	result := -1

	for i := 0; i <= len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}
	return result
}
