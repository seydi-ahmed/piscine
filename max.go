package piscine

func Max(a []int) int {
	max := a[0]
	if len(a) == 0 {
		return 0
	} else {
		for i := 0; i < len(a); i++ {
			if max < a[i] {
				max = a[i]
			}
		}
	}
	return max
}
