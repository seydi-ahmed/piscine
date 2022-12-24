package piscine

func ListPushBack(l *List, data interface{}) {
	n := NodeL{Data: data, Next: nil}
	if l.Head == nil {
		l.Head = &n
		l.Tail = &n
	} else {
		l.Tail.Next = &n
		l.Tail = &n
	}
}
