package piscine

func SplitWhiteSpaces(s string) []string {
	runes := []rune(s)
	result := []string{}
	index := 0
	inWord := true
	for i := 0; i < len(runes); i++ {
		if runes[i] == '\n' || runes[i] == ' ' || runes[i] == '\t' {
			if inWord {
				index++
			}
			inWord = false
		} else {
			inWord = true
			if len(result) > index {
				result[index] += string(runes[i])
			} else {
				result = append(result, string(runes[i]))
			}
		}
	}
	return result
}
