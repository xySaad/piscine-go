package piscine

func Compact(ptr *[]string) int {
	myslice := *ptr
	compactedSlice := []string{}

	for i := 0; i < len(myslice); i++ {
		if myslice[i] != "" {
			compactedSlice = append(compactedSlice, myslice[i])
		}
	}
	*ptr = compactedSlice

	return len(compactedSlice)
}
