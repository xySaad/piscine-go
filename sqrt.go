package piscine

func Sqrt(nb int) int {
	var result int
	for i := range nb {
		if i*i == nb {
			result = i
		}
	}
	return result
}
