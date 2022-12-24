package piscine

func Index(s string, toFind string) int {
	for i := 0; i < len(s)-len(toFind); i++ {
		if toFind == s[i:i+len(toFind)] {
			return i
		}
	}
	return -1
}
