package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	}
	arrayLength := max - min
	newArray := make([]int, arrayLength)

	for i := range arrayLength {
		newArray[i] = i + min
	}
	return newArray
}
