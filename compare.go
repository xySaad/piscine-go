package piscine

func Compare(a, b string) int {
	var result int
	if a < b {
		result = -1
	} else if a == b {
		result = 0
	} else if a > b {
		result = 1
	}
	return result
}
