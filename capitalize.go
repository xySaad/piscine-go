package piscine

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
