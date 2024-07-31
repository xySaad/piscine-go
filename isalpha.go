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
