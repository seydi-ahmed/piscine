package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	for _, lettre := range arguments {
		lettreRune := []rune(lettre)
		for i := 0; i < len(lettreRune); i++ {
			z01.PrintRune(lettreRune[i])
		}
		z01.PrintRune('\n')
	}
}
