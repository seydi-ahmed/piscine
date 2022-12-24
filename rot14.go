package piscine

func pos(lettre rune) rune {
	if lettre == 'a' {
		return 'o'
	}
	if lettre == 'A' {
		return 'O'
	}
	if lettre == 'b' {
		return 'p'
	}
	if lettre == 'B' {
		return 'P'
	}
	if lettre == 'c' {
		return 'q'
	}
	if lettre == 'C' {
		return 'Q'
	}
	if lettre == 'd' {
		return 'r'
	}
	if lettre == 'D' {
		return 'R'
	}
	if lettre == 'e' {
		return 's'
	}
	if lettre == 'E' {
		return 'S'
	}
	if lettre == 'f' {
		return 't'
	}
	if lettre == 'F' {
		return 'T'
	}
	if lettre == 'g' {
		return 'u'
	}
	if lettre == 'G' {
		return 'U'
	}
	if lettre == 'h' {
		return 'v'
	}
	if lettre == 'H' {
		return 'V'
	}
	if lettre == 'i' {
		return 'w'
	}
	if lettre == 'I' {
		return 'W'
	}
	if lettre == 'j' {
		return 'x'
	}
	if lettre == 'J' {
		return 'X'
	}
	if lettre == 'k' {
		return 'y'
	}
	if lettre == 'K' {
		return 'Y'
	}
	if lettre == 'l' {
		return 'z'
	}
	if lettre == 'L' {
		return 'Z'
	}
	if lettre == 'm' {
		return 'a'
	}
	if lettre == 'M' {
		return 'A'
	}
	if lettre == 'n' {
		return 'b'
	}
	if lettre == 'N' {
		return 'B'
	}
	if lettre == 'o' {
		return 'c'
	}
	if lettre == 'O' {
		return 'C'
	}
	if lettre == 'p' {
		return 'd'
	}
	if lettre == 'P' {
		return 'D'
	}
	if lettre == 'q' {
		return 'e'
	}
	if lettre == 'Q' {
		return 'E'
	}
	if lettre == 'r' {
		return 'f'
	}
	if lettre == 'R' {
		return 'F'
	}
	if lettre == 's' {
		return 'g'
	}
	if lettre == 'S' {
		return 'G'
	}
	if lettre == 't' {
		return 'h'
	}
	if lettre == 'T' {
		return 'H'
	}
	if lettre == 'u' {
		return 'i'
	}
	if lettre == 'U' {
		return 'I'
	}
	if lettre == 'v' {
		return 'j'
	}
	if lettre == 'V' {
		return 'J'
	}
	if lettre == 'w' {
		return 'k'
	}
	if lettre == 'W' {
		return 'K'
	}
	if lettre == 'x' {
		return 'l'
	}
	if lettre == 'X' {
		return 'L'
	}
	if lettre == 'y' {
		return 'm'
	}
	if lettre == 'Y' {
		return 'M'
	}
	if lettre == 'z' {
		return 'n'
	}
	if lettre == 'Z' {
		return 'N'
	}
	return 'i'
}

func Rot14(s string) string {
	s1 := []rune(s)
	s2 := ""
	for i := 0; i < len(s1); i++ {
		if s1[i] >= 'a' && s1[i] <= 'z' {
			s2 = s2 + string(pos(s1[i]))
		} else if s1[i] >= 'A' && s1[i] <= 'Z' {
			s2 = s2 + string(pos(s1[i]))
		} else {
			s2 = s2 + string(s1[i])
		}
	}
	return s2
}
