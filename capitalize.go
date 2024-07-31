package piscine

func Capitalize(s string) string {
	shouldCapitalize := true
	result := []rune(s)

	for i := 0; i < len(result); i++ {
		if IsAlpha(string(s[i])) {
			if shouldCapitalize {
				if IsLower(string(s[i])) {
					result[i] = rune(s[i] - 32)
				} else {
					result[i] = rune(s[i])
				}
			} else if IsUpper(string(s[i])) {
				result[i] = rune(s[i] + 32)
			} else {
				result[i] = rune(s[i])
			}
			shouldCapitalize = false
		} else {
			result[i] = rune(s[i])
			shouldCapitalize = true
		}
	}
	return string(result)
}
