package piscine

func StrRev(s string) string {
	stringLength := 0

	for range s {
		stringLength++
	}
	reversedString := ""
	for i := 1; i <= stringLength; i++ {
		reversedString += string(s[stringLength-i])
	}
	return reversedString
}
