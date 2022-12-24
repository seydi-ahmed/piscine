package piscine

func aZero(s string) string {
	return ""
}

func nbrSpace(s string) int {
	c := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			c++
		}
	}
	return c
}

func viderTable(tab []string) []string {
	var tab1 []string
	for i := 0; i < len(tab); i++ {
		if tab[i] != "" {
			tab1 = append(tab1, tab[i])
		}
	}
	return tab1
}

func SplitWhiteSpaces(s string) []string {
	count := nbrSpace(s)
	space := 0
	var test []string
	avance := 0

	temoin := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\n' && s[i] != '\t' {
			temoin += string(s[i])
		} else {
			test = append(test, temoin)
			temoin = aZero(temoin)
			space++
			if count == space {
				break
			}
		}
		avance++
	}
	temoin = aZero(temoin)
	for i := avance + 1; i < len(s); i++ {
		temoin += string(s[i])
	}
	test = append(test, temoin)
	test = viderTable(test)
	return test
}
