package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	l1 := l.Head
	var newlist List

	for l1 != nil {
		if l1.Data != data_ref {
			ListPushBack(&newlist, l1.Data)
		}
		l1 = l1.Next
	}
	l.Head = newlist.Head
	l.Tail = newlist.Tail
}
