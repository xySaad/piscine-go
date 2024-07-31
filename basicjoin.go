package piscine

func BasicJoin(elems []string) string {
	newString := ""

	for i := range elems {
		newString += elems[i]
	}

	return newString
}
