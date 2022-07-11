package main

import (
	"github.com/01-edu/z01"

	"os"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func even(str int) int {
	Arg := os.Args[1:]
	a := 0
	for i := 0; i <= len(Arg); i++ {
		a = i
	}
	return a
}
func isEven(nbr int) bool {
	if even(nbr) == 1 {
		return false
	} else {
		return true
	}
}

func main() {

	lengthofArg := 0
	if isEven(lengthofArg) == true {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
