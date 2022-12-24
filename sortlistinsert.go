package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	n := &NodeI{Data: data_ref, Next: nil}

	if l == nil || l.Data >= data_ref {
		n.Next = l
		return n
	} else {
		l1 := l
		for l1.Next != nil && l1.Next.Data < data_ref {
			l1 = l1.Next
		}

		n.Next = l1.Next
		l1.Next = n
	}
	return l
}
