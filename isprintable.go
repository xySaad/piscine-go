package piscine

func IsPrintable(s string) bool {
	for ch := range s {
		if s[ch] < 32 || s[ch] > 126 {
			return false
		}
	}
	return true
}
