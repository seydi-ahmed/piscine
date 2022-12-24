package piscine

func SortIntegerTable(table []int) {
	for i := range table {
		for j := range table {
			if table[i] < table[j] {
				Swap(&table[j], &table[i])
			}
		}
	}
}
