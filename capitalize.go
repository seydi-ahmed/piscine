package piscine

func Capitalize(s string) string {
	s1 := []rune(s)
	isFirstLetterWord := false
	isOnWord := false
	for i := 0; i < len(s); i++ {
		if IsAlpha(string(s[i])) {
			if isFirstLetterWord && isOnWord {
				isFirstLetterWord = false
			} else if !isFirstLetterWord && !isOnWord {
				isFirstLetterWord = true
			}
			isOnWord = true
		} else {
			isOnWord = false
		}
		if isFirstLetterWord {
			if !IsNumeric(string(s[i])) && IsLower(string(s[i])) {
				s1[i] = rune(ToUpper(string(s1[i]))[0])
			}
			isFirstLetterWord = false
		} else if isOnWord {
			s1[i] = rune(ToLower(string(s1[i]))[0])
		}
	}
	return string(s1)
}
