package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil {
		if n2 == nil {
			return nil
		} else {
			return n2
		}
	}

	if n2 == nil {
		if n1 == nil {
			return nil
		} else {
			return n1
		}
	}

	var n *NodeI = nil

	for n1 != nil {
		n = ajout(n, n1.Data)
		n1 = n1.Next
	}

	for n2 != nil {
		n = ajout(n, n2.Data)
		n2 = n2.Next
	}

	ListSort(n)
	return n
}

func ajout(n *NodeI, data int) *NodeI {
	noeud := &NodeI{Data: data, Next: nil}
	n1 := n
	if n1 == nil {
		return noeud
	}
	for n1.Next != nil {
		n1 = n1.Next
	}
	n1.Next = noeud
	return n
}
