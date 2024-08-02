package piscine

func SplitWhiteSpaces(s string) []string {
	runes := []rune(s)
	result := []string{}
	index := 0
	for i := 0; i < len(runes); i++ {
		if runes[i] == '\n' || runes[i] == ' ' || runes[i] == '\t' {
			for j := 0; runes[i] == '\n' || runes[i] == ' ' || runes[i] == '\t'; j++ {
				if j > 0 {
					i++
				}
			}
		}
		if len(result) > index {
			result[index] += string(runes[i])
		} else {
			result = append(result, string(runes[i]))
		}
	}
	return result
}
