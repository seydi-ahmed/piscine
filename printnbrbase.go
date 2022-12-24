package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '-' || base[i] == '+' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
	}
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}
	if nbr < 0 {
		z01.PrintRune('-')
	}
	res := ""
	ra := []rune(base)

	for ; nbr != 0; nbr /= len(base) {
		index := nbr % len(base)
		if index < 0 {
			index = -index
		}
		res = string(ra[index]) + res
	}
	for _, v := range res {
		z01.PrintRune(v)
	}
}
