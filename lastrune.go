package piscine

func LastRune(s string) rune {
	runeSlice := []rune(s)

	return runeSlice[len(s)-1]
}
