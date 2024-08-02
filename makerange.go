package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	}
	arrayLength := max - min
	newArray := make([]int, arrayLength)

	for i := 0; i <= arrayLength-1; i++ {
		newArray[i] = i + min
	}
	return newArray
}
