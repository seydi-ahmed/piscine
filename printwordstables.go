package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i := 0; i < len(a); i++ {
		temoin := a[i]
		for i := 0; i < len(temoin); i++ {
			z01.PrintRune(rune(temoin[i]))
		}
		z01.PrintRune('\n')
	}
}
