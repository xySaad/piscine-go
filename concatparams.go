package piscine

func ConcatParams(args []string) string {
	result := ""
	seperator := "\n"

	for i := 0; i < len(args); i++ {
		result += args[i]
		if i < len(args)-1 {
			result += seperator
		}
	}
	return result
}
