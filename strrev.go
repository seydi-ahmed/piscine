package piscine

func StrRev(s string) string {
	runes := []rune(s)
	revrune := []rune{}
	for i := len(runes) - 1; i >= 0; i-- {
		revrune = append(revrune, runes[i])
	}
	str := string(revrune)
	return str
}
