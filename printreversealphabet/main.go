package main

import "github.com/01-edu/z01"

func main() {
	c := 'a'
	t := 'z'

	for i := t; i >= c; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
