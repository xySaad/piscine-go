package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}

	var result interface{}

	current := l.Head

	for current.Next != nil {
		current = current.Next
	}

	result = current.Data

	return result
}
