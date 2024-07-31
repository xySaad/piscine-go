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
