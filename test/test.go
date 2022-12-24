package main

import "fmt"

func main() {
	s1 := "hello"
	s2 := "hbujnnelknknlhbvhbodbfv"
	k := 0
	text := ""
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] && j > k {
				k = j
				text = text + string(s2[j])
				break
			}
		}
	}
	if text == s1 {
		fmt.Println(text)
	}
}
