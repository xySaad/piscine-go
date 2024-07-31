package piscine

func TrimAtoi(s string) int {
	sign := 1
	result := 0

	for i := range s {
		if s[i] <= '9' && s[i] >= '0' {
			break
		}
		if s[i] == '-' {
			sign = -1
			break
		}
		if s[i] == '+' {
			sign = 1
			break
		}
	}
	for i := range s {
		if s[i] <= '9' && s[i] >= '0' {
			result *= 10
			result += int(s[i] - '0')
		}
	}
	return result * sign
}
