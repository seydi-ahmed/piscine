package piscine

func BasicAtoi(s string) int {
	nb := 0
	f := 1
	for i := len(s) - 1; i >= 0; i-- {
		nb += (int(rune(s[i])) - 48) * f
		f = f * 10
	}
	return nb
}
