package piscine

import "github.com/01-edu/z01"

func FirstRune(s string) rune {
	sentence := []rune(s)

	for index, letter := range sentence {
		if index == 0 {
			z01.PrintRune(letter)
		}
	}
	return 0
}
