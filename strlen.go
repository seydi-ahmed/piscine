package piscine

func StrLen(s string) int {
	l := 0
	a := []rune(s)
	for index := range a {
		l = index + 1
	}
	return l
}
