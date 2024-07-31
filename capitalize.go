package piscine

func IsAlpha(s string) bool {
	if s == "" {
		return true
	}

	for i := range s {
		if s[i] < 'A' && s[i] > '9' {
			return false
		}
		if s[i] < 'a' && s[i] > 'Z' {
			return false
		}
		if s[i] < '0' || s[i] > 'z' {
			return false
		}
	}

	return true
}

func IsLower(s string) bool {
	for i := range s {
		if s[i] < 'a' || s[i] > 'z' {
			return false
		}
	}
	return true
}

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
