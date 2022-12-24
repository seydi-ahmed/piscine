package piscine

func ListReverse(l *List) {
	var newList List

	for l.Head != nil {
		ListPushFront(&newList, l.Head.Data)
		l.Head = l.Head.Next
	}
	// ListClear(l)
	l.Head = newList.Head

	l.Tail = newList.Tail
}
