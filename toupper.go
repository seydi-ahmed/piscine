package piscine

func ToUpper(s string) string {
	s1 := []rune(s)
	for i := 0; i < len(s1); i++ {
		for k, j := 'a', 'A'; k <= 'z'; k++ {
			if s1[i] == k {
				s1[i] = j
			}
			j++
		}
	}
	return string(s1)
}
