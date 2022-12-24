package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	l1 := l

	if l == nil {
		return nil
	}
	if l.Next == nil {
		return l
	}
	for l1 != nil {
		l2 := l1.Next
		for l2 != nil {
			if l1.Data > l2.Data {
				l2.Data, l1.Data = l1.Data, l2.Data
			}
			l2 = l2.Next
		}
		l1 = l1.Next
	}

	return l
}
