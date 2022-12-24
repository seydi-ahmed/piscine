package piscine

func Atoi(s string) int {
	nb := 0
	c := 0
	signe := 1

	for i := 0; i < len(s); i++ {
		c = int(s[i]) - 48
		if i == 0 {
			if c == -5 {
				continue
			} else if c == -3 {
				signe = -1
			} else {
				nb = c
			}
		} else {
			if c >= 0 && c <= 9 {
				nb = nb*10 + c
			} else {
				nb = 0
				break
			}
		}
	}

	nb = nb * signe
	return nb
}
