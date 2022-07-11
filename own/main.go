package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	SwitchCase(arg)
}

func SwitchCase(arg []string) {
	for ind := range arg {
		if ind > 1 {
			break
		} else {
			for a := 0; a < len(arg); a++ {
				for _, v := range arg[ind] {
					if v == 'a' || v == 'A' {
						z01.PrintRune(v)
					} else if v == 'b' || v == 'B' {
						for i := 2; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'c' || v == 'C' {
						for i := 3; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'd' || v == 'D' {
						for i := 4; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'e' || v == 'E' {
						for i := 5; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'f' || v == 'F' {
						for i := 6; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'g' || v == 'G' {
						for i := 7; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'h' || v == 'H' {
						for i := 8; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'i' || v == 'I' {
						for i := 9; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'j' || v == 'J' {
						for i := 10; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'k' || v == 'K' {
						for i := 11; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'l' || v == 'L' {
						for i := 12; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'm' || v == 'M' {
						for i := 13; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'n' || v == 'N' {
						for i := 14; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'o' || v == 'O' {
						for i := 15; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'p' || v == 'P' {
						for i := 16; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'q' || v == 'Q' {
						for i := 17; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'r' || v == 'R' {
						for i := 18; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 's' || v == 'S' {
						for i := 19; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 't' || v == 'T' {
						for i := 20; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'u' || v == 'U' {
						for i := 21; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'v' || v == 'V' {
						for i := 22; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'w' || v == 'W' {
						for i := 23; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'x' || v == 'X' {
						for i := 24; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'y' || v == 'Y' {
						for i := 25; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else if v == 'z' || v == 'Z' {
						for i := 26; i >= 1; i-- {
							z01.PrintRune(v)
						}
					} else {
						z01.PrintRune(v)
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
