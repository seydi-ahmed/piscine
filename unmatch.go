package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		c := 0
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				c++
			}
		}
		if c%2 != 0 {
			return a[i]
		}
	}
	return -1
}
