package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	l1 := l.Head

	for l1 != nil {
		if comp(l1.Data, ref) {
			return &l1.Data
		}
		l1 = l1.Next
	}
	return nil
}
