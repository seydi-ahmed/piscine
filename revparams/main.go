package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for i := len(args) - 1; i >= 0; i-- {
		temoin := args[i]
		for j := 0; j < len(temoin); j++ {
			z01.PrintRune(rune(temoin[j]))
		}
		z01.PrintRune('\n')
	}
}
