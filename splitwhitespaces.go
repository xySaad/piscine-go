package piscine

func SplitWhiteSpaces(s string) []string {
	runes := []rune(s)
	result := []string{}
	index := 0
	for i := 0; i < len(runes); i++ {
		if runes[i] == '\n' || runes[i] == ' ' {
			index++
		} else {
			if len(result) > index {
				result[index] += string(runes[i])
			} else {
				result = append(result, string(runes[i]))
			}
		}
	}
	return result
}
