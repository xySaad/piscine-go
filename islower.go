package piscine

func IsLower(s string) bool {
	for i := range s {
		if s[i] < 'a' || s[i] > 'z' {
			return false
		}
	}
	return true
}
