package piscine

func StringToIntSlice(str string) []int {
	result := []int(nil)
	StringAsRunes := []rune(str)
	for i := range StringAsRunes {
		result = append(result, int(StringAsRunes[i]))
	}
	return result
}
