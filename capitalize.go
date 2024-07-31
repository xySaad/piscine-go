package piscine

func ToUpper(s string) string {
	stringUpperCase := ""

	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			stringUpperCase += string(s[i] - 32)
		} else {
			stringUpperCase += string(s[i])
		}
	}
	return stringUpperCase
}

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

func Capitalize(s string) string {
	shouldCapitalize := true
	newString := ""

	for i := range s {
		if IsAlpha(string(s[i])) {
			if shouldCapitalize {
				newString += ToUpper(string(s[i]))
			} else {
				newString += string(s[i])
			}
			shouldCapitalize = false
		} else {
			newString += string(s[i])
			shouldCapitalize = true
		}
	}
	return newString
}
