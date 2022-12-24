package piscine

func BasicAtoi2(s string) int {
	runes := []rune(s)
	nb := 0
	f := len(s)
	for i := 0; i < f; i++ {
		if runes[i] < '0' || runes[i] > '9' {
			return 0
		} else {
			nb *= 10
			nb += int(runes[i]) - '0'
		}
	}
	return nb
}
