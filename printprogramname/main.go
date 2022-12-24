package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, w := range os.Args[0] {
		if w != '.' && w != '/' {
			z01.PrintRune(w)
		}
	}
	z01.PrintRune('\n')
}
