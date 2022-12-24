package piscine

func aBase2(nb int) []int {
	var tab []int
	for nb != 0 {
		tab = append(tab, nb%2)
		nb = nb / 2
	}
	return tab
}

func ActiveBits(n int) int {
	tab := aBase2(n)
	c := 0
	for i := 0; i < len(tab); i++ {
		if tab[i] == 1 {
			c++
		}
	}
	return c
}
