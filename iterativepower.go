package piscine

func IterativePower(nb int, power int) int {
	if power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	}
	b := 1
	for a := power; a >= 1; a-- {
		b *= nb
	}
	return b
}
