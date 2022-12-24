package piscine

func Split(s, sep string) []string {
	longueur := 0
	for i := range sep {
		longueur = i + 1
	}

	longueur2 := 0
	for i := range s {
		longueur2 = i + 1
	}

	for i := 0; i < longueur2-longueur; i++ {
		if s[i:i+longueur] == sep {
			s = s[:i] + " " + s[i+longueur:]
			longueur2 = longueur2 - longueur
		}
	}
	t := SplitWhiteSpaces(s)
	return t
}
