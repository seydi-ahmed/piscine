package piscine

func CollatzCountdown(start int) int {
	nb := 0
	start1 := start

	if start1 <= 0 {
		return -1
	}
	for start1 != 1 {
		if start1%2 == 0 {
			start1 = start1 / 2
		} else {
			start1 = start1*3 + 1
		}
		nb++
	}
	return nb
}
