package piscine

func ListClear(l *List) {
	if l.Head == nil {
	} else {
		l.Head = nil
		l.Tail = nil
	}
}
