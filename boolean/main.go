package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	s1 := []rune(s)
	for i := 0; i < len(s1); i++ {
		z01.PrintRune(s1[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	return nbr%2 == 0
}

func main() {
	arguments := os.Args[1:]
	lengthOfArg := len(arguments)

	if isEven(lengthOfArg) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
