package piscine

func ListSize(l *List) int {
	counter := 0

	for l != nil {
		l.Head = l.Head.Next
		counter++
	}
	return counter
}
