package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	c := 0
	if l.Head == nil {
		return 0
	} else {
		for l.Head != nil {
			c++
			l.Head = l.Head.Next
		}
	}
	return c
}
