package piscine

func ToLower(s string) string {
	stringLowerCase := ""

	for i := range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			stringLowerCase += string(s[i] + 32)
		} else {
			stringLowerCase += string(s[i])
		}
	}
	return stringLowerCase
}
