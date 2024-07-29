package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table)-1; i++ {
		for i := 0; i < len(table)-1; i++ {
			prev := table[i]
			next := table[i+1]
			if prev > next {
				table[i] = next
				table[i+1] = prev
			}
		}

	}
}
