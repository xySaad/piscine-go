package piscine

func CountIf(f func(string) bool, tab []string) int {
	for i := 0; i < len(tab); i++ {
		if f(tab[i]) {
			return i + 1
		}
	}
	return 0
}
