package piscine

import "fmt"

func FirstRune(s string) rune {
	sentence := []rune(s)

	for index, letter := range sentence {
		if index == 0 {
			fmt.Printf("%c", letter)
		}
	}
	return 0
}
