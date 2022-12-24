package piscine

func IsUpper(s string) bool {
	s1 := []rune(s)
	count := 0
	for i := 0; i < len(s1); i++ {
		for j := 'A'; j <= 'Z'; j++ {
			if s1[i] == j {
				count++
			}
		}
	}
	if count == len(s1) {
		return true
	} else {
		return false
	}
}
