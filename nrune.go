package piscine

func NRune(s string, n int) rune {
	s1 := []rune(s)
	taille := len(s1)
	if n > 0 && n <= taille {
		return s1[n-1]
	}
	return rune(0)
}
