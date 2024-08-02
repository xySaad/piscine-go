package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return []int{}
	}

	newArray := []int{}

	for i := min; i < max; i++ {
		newArray = append(newArray, i)
	}
	return newArray
}
