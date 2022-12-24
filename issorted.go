package piscine

func Comp(a, b int) int {
	if a-b > 0 {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func IsSorted(f func(a, b int) int, a []int) bool {
	croissant := true
	decroissant := true

	for i := range a {
		if i < len(a)-1 {
			{
				if Comp(a[i], a[i+1]) < 0 {
					decroissant = false
				}
				if Comp(a[i], a[i+1]) > 0 {
					croissant = false
				}
			}
		}
	}
	return croissant || decroissant
}
