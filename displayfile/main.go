package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]

	if len(arguments) == 0 {
		fmt.Println("File name missing")
		return
	}

	if len(arguments) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	count := 0
	for _, lettre := range arguments {
		for i := 0; i < len(lettre); i++ {
			if lettre[i] == '.' {
				count++
			}
		}
	}
	if count > 1 {
		fmt.Println("Too many arguments")
		return
	}
	if count == 1 {
		fmt.Println("Almost there!!")
	}
}
