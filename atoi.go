package piscine

func Atoi(s string) int {
	var numberAsInteger int = 0
	var sign int = 1
	for i := range s {
		var num int = 0
		if s[i] >= '0' && s[i] <= '9' {
			num = int(s[i] - '0')
		} else if i == 0 {
			switch s[i] {
			case '-':
				sign = -1
			case '+':
				sign = 1
			default:
				return 0
			}
		} else {
			return 0
		}
		numberAsInteger = numberAsInteger*10 + num
	}
	return numberAsInteger * sign
}
