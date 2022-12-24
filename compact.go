package piscine

func isNul(s string) bool {
	return s == ""
}

func Compact(ptr *[]string) int {
	c := 0
	ptr1 := *ptr
	for i := 0; i < len(ptr1); i++ {
		for j := i + 1; j < len(ptr1); j++ {
			if isNul(ptr1[i]) {
				ptr1[i] = ptr1[j]
				ptr1[j] = ""
			}
		}
	}
	ptr2 := []string{}
	for i := 0; i < len(ptr1); i++ {
		if !(isNul(ptr1[i])) {
			ptr2 = append(ptr2, ptr1[i])
			c++
		}
	}
	*ptr = ptr2
	return c
}
