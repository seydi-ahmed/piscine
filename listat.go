package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	c := 0
	if l == nil {
		return nil
	} else {
		for l != nil {
			if pos == c {
				return l
			}
			l = l.Next
			c++
		}
	}
	return nil
}
