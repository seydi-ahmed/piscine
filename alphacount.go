package piscine

func AlphaCount(s string) int {
	count := 0
	alphabet := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	alphaBet := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	s1 := []rune(s)
	taille := len(s1)

	for i := 0; i < taille; i++ {
		for j := range alphabet {
			if s1[i] == alphabet[j] {
				count++
			}
		}
	}

	for i := 0; i < taille; i++ {
		for k := range alphabet {
			if s1[i] == alphaBet[k] {
				count++
			}
		}
	}

	return count
}
