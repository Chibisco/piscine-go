package piscine

func RecursiveFactorial(nb int) int {
	a := 1
	if nb == 0 || nb == 1 {
		return 1
	} else if a < nb {
		return nb * RecursiveFactorial(nb-1)
	}
	return 0
}
