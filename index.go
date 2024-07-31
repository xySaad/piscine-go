package piscine

func Index(s string, toFind string) int {
	var result int

	for i := range s {
		if toFind[0] == s[i] {
			for j := 0; j < len(toFind); j++ {
				if s[i+j] != toFind[j] {
					return -1
				}
			}
			result++
		}
	}
	if result == 0 {
		return -1
	}
	return result
}
