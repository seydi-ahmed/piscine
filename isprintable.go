package piscine

func charIsPrintable(r rune) bool {
	if r >= 0 && r <= 31 {
		return true
	}
	return false
}

func IsPrintable(s string) bool {
	s1 := []rune(s)
	for _, i := range s1 {
		if charIsPrintable(i) {
			return false
		}
	}
	return true
}
