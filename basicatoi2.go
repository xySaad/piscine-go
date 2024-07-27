package piscine

func BasicAtoi2(s string) int {
	var numberAsInteger int = 0
	for i := range s {
		var num int = 0
		if s[i] >= '0' && s[i] < '9' {
			num = int(s[i] - '0')
		}
		numberAsInteger = numberAsInteger*10 + num
	}
	return numberAsInteger
}
