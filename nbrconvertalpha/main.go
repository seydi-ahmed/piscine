package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) >= 2 {
		arg1 := os.Args[1:]
		isUpper := false
		if arg1[0] == "--upper" {
			isUpper = true
			arg1 = arg1[1:]
		}
		for _, arg := range arg1 {
			nb := 0
			nb = BasicAtoi(arg)
			if nb >= 1 && nb <= 26 {
				if !isUpper {
					z01.PrintRune(rune(nb + 96))
				} else {
					z01.PrintRune(rune(nb + 64))
				}
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func BasicAtoi(s string) int {
	nb := 0
	f := 1
	for i := len(s) - 1; i >= 0; i-- {
		nb += (int(rune(s[i])) - 48) * f
		f = f * 10
	}
	return nb
}
