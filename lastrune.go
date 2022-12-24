package piscine

func LastRune(s string) rune {
	s1 := []rune(s)
	tailleS1 := len(s1)
	return s1[tailleS1-1]
}
