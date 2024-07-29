package piscine

func Sqrt(nb int) int {
	var result int
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			result = i
		}
	}
	return result
}
