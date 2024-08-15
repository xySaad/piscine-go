package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	var result *NodeL
	counter := 0
	current := l
	for counter != pos {
		if current.Next == nil {
			return nil
		}
		current = current.Next
		counter++
	}

	result = current
	return result
}
