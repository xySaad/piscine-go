package piscine

func NRune(s string, n int) rune {
	var requestedChar rune
	if n <= len(s) && n > 0 {
		stringChars := []rune(s)
		requestedChar = stringChars[n-1]
	} else {
		return 0
	}

	return requestedChar
}
