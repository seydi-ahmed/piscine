package main

import (
	"os"

	"github.com/01-edu/z01"
)

func order(s string) {
	var a [5000]int
	for _, c := range s {
		a[int(c)]++
	}
	for i, c := range a {
		for c > 0 {
			z01.PrintRune(rune(i))
			c--
		}
	}
	z01.PrintRune('\n')
}

func insert(str1 string, str2 string) string {
	return str1 + str2
}

func help() {
	insert := "--insert"
	insert1 := []rune(insert)
	ii := "  -i"
	ii1 := []rune(ii)
	this_flag := "	 This flag inserts the string into the string passed as argument."
	this_flag1 := []rune(this_flag)
	order := "--order"
	order1 := []rune(order)
	oo := "  -o"
	oo1 := []rune(oo)
	flag_will := "	 This flag will behave like a boolean, if it is called it will order the argument."
	flag_will1 := []rune(flag_will)

	for i := 0; i < len(insert1); i++ {
		z01.PrintRune(insert1[i])
	}
	z01.PrintRune('\n')
	for i := 0; i < len(ii1); i++ {
		z01.PrintRune(ii1[i])
	}
	z01.PrintRune('\n')
	for i := 0; i < len(this_flag1); i++ {
		z01.PrintRune(this_flag1[i])
	}
	z01.PrintRune('\n')
	for i := 0; i < len(order1); i++ {
		z01.PrintRune(order1[i])
	}
	z01.PrintRune('\n')
	for i := 0; i < len(oo1); i++ {
		z01.PrintRune(oo1[i])
	}
	z01.PrintRune('\n')
	for i := 0; i < len(flag_will1); i++ {
		z01.PrintRune(flag_will1[i])
	}
	z01.PrintRune('\n')
}

func main() {
	myArr := os.Args[1:]
	ln := -1

	for i := range myArr {
		ln = i
	}
	if ln != -1 {
		if myArr[0] == "-h" || myArr[0] == "--help" {
			help()
		} else if myArr[0] == "--order" || myArr[0] == "-o" {
			order(myArr[1])
		} else if myArr[0][0:3] == "--i" || myArr[0][0:2] == "-i" {
			if ln < 2 {
				if myArr[0][0:3] == "--i" {
					temoin := insert(myArr[1], myArr[0][9:])
					for i := 0; i < len(temoin); i++ {
						z01.PrintRune(rune(temoin[i]))
					}
					z01.PrintRune('\n')
				} else {
					temoin := insert(myArr[1], myArr[0][3:])
					for i := 0; i < len(temoin); i++ {
						z01.PrintRune(rune(temoin[i]))
					}
					z01.PrintRune('\n')
				}
			} else {
				myStr := ""
				for i := 1; i <= ln; i++ {
					if myArr[i] != "-o" && myArr[i] != "--order" {
						myStr = insert(myStr, myArr[i])
					}
				}

				if myArr[0][0:3] == "--i" {
					myStr = insert(myStr, myArr[0][9:])
				} else {
					myStr = insert(myStr, myArr[0][3:])
				}
				order(myStr)
			}
		} else {
			temoin := myArr[0]
			for i := 0; i < len(temoin); i++ {
				z01.PrintRune(rune(temoin[i]))
			}
			z01.PrintRune('\n')
		}
	} else {
		help()
	}
}
